package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"

	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

func NewMqttProtocol(conn connection.ConnectionInterface) *MqttProtocol {
	return &MqttProtocol{
		conn:   conn,
		logger: logger.NewLogger("Mqtt"),
	}
}

func (prot *MqttProtocol) UpdateLogger(logger *logger.Logger) {
	prot.logger = logger
}

func (prot *MqttProtocol) connectUnPack() (*ResponseConnect, error) {
	var response ResponseConnect
	data, err := prot.isMqttCmd(CONNECT)
	if err != nil {
		return nil, err
	}
	msg := data[9:]
	if len(msg) < 3 {
		return nil, fmt.Errorf("length data is < 3")
	}
	connectFlag := msg[0]
	if connectFlag&0b1 == 0b1 {
		return nil, fmt.Errorf("Reserved flag on connectFlag isn't 0")
	}
	if connectFlag&0b10 != 0b10 { //CleanSession is 0
		return nil, nil
	}
	var hasWillFlag = false
	//TODO processar o will flags

	var hasUser = false
	if connectFlag&0b10000000 == 0b10000000 { //has username is 1
		hasUser = true
	}
	var hasPassword = false
	if connectFlag&0b1000000 == 0b1000000 { //has hasPassword is 1
		hasPassword = true
		if !hasUser {
			hasPassword = false
		}
	}

	var keepAlive int16
	err = binary.Read(bytes.NewReader(msg[1:3]), binary.BigEndian, &keepAlive)
	if err != nil {
		return &response, fmt.Errorf("Erro ao ler bytes:%s", err)
	}
	var clientId, username, password string
	var lengthWill, lengthUser, lengthClientId, lengthPayload int

	payload := msg[3:]
	prot.logger.Debug("Payload: %v", payload)
	lengthPayload = len(payload)
	if lengthPayload < 2 {
		//TODO criar client Id random se nao tiver sido passado
		//("Nao forneceu ClientID")
		clientId = "ClientId_Random"
		response = ResponseConnect{
			Id:        clientId,
			Timeout:   1,
			Clean:     true,
			KeepAlive: keepAlive,
			Username:  "",
			Password:  "",
		}
		return &response, nil
	}
	lengthClientId = int(payload[0])<<8 + int(payload[1])
	payload = payload[2:]
	clientId = string(payload[:lengthClientId])
	if len(payload) > lengthClientId {
		payload = payload[lengthClientId:]
	}

	//TODO will msg
	if hasWillFlag {
		prot.logger.Debug("hasWillFlag")
		//lengthWill = 0
		payload = payload[lengthWill:]
	}
	if hasUser {
		prot.logger.Debug("hasUser")
		lengthUser = int(payload[0])<<8 + int(payload[1]) //TODO calc length
		payload = payload[2:]
		username = string(payload[:lengthUser])
		payload = payload[lengthUser:]
	}
	if hasPassword {
		prot.logger.Debug("hasPassword")
		lengthPassword := int(payload[0])<<8 + int(payload[1])
		payload = payload[2:]
		password = string(payload[:lengthPassword])
	}
	response = ResponseConnect{
		Id:        clientId,
		Timeout:   1,
		Clean:     true,
		KeepAlive: keepAlive,
		Username:  username,
		Password:  password,
	}
	return &response, nil
}

func (prot *MqttProtocol) connAck(sessionCfg *ResponseConnect, resp ConnectReturnCode) error {
	response := []byte{byte(COMMAND_CONNACK), 0b10}
	if sessionCfg.Clean {
		response = append(response, []byte{0x00}...)
	} else {
		response = append(response, []byte{0x01}...)
	}
	response = append(response, []byte{byte(resp)}...) //TODO response
	prot.logger.Debug("connAck: %v", response)
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func (prot *MqttProtocol) publishUnPack(data []byte) (*ResponsePublish, error) {
	prot.logger.Debug("publishUnPack")
	var response ResponsePublish
	response.dutFlag = (data[0]&0b1000 == 0b1000)
	response.Qos = int((data[0] & 0b110) >> 1)
	response.Retained = (data[0]&0b1 == 0b1)
	variableHeader := data[2:]
	if len(variableHeader) < 3 {
		return nil, fmt.Errorf("tamanho do payload(%d) sem topic: %v", len(variableHeader), variableHeader)
	}
	lengthTopic := int(variableHeader[0])<<8 + int(variableHeader[1])
	variableHeader = variableHeader[2:]
	response.Topic = string(variableHeader[:lengthTopic])
	variableHeader = variableHeader[lengthTopic:]
	if response.Qos > 0 {
		response.Identifier = variableHeader[:2]
		variableHeader = variableHeader[2:]
	}
	payload := variableHeader
	response.Payload = payload
	return &response, nil
}
func (prot *MqttProtocol) pubAck(pubCfg *ResponsePublish) error {
	prot.logger.Debug("pubAck")
	response := []byte{byte(COMMAND_PUBACK), 0b10}
	response = append(response, pubCfg.Identifier...)
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
func (prot *MqttProtocol) pubRec(pubCfg *ResponsePublish) error {
	prot.logger.Debug("pubRec")
	response := []byte{byte(COMMAND_PUBREC), 0b10}
	response = append(response, pubCfg.Identifier...)
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
func (prot *MqttProtocol) unpackPubRel(data []byte, packetIdentifier *[]byte) error {
	prot.logger.Debug("unpackPubRel")
	if data[0] != byte(COMMAND_PUBREL) {
		return fmt.Errorf("Byte command isn't exactly COMMAND_PUBREL")
	}
	length := len(data[1:])
	if length < 2 {
		return fmt.Errorf("length < 2")
	}
	if !bytes.Equal(data[2:], *packetIdentifier) {
		return fmt.Errorf("incorrect packet identifier")
	}
	return nil
}
func (prot *MqttProtocol) pubComp(packetIdentifier *[]byte) error {
	prot.logger.Debug("pubRec")
	response := []byte{byte(COMMAND_PUBCOMP), 0b10}
	response = append(response, *packetIdentifier...)
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
