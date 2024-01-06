package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
)

func NewMqttProtocol(conn connection.ConnectionInterface) *MqttProtocol {
	return &MqttProtocol{
		conn: conn,
	}
}

func (prot *MqttProtocol) IsValidMqttCmd() (*Command, []byte, error) {
	data, err := prot.conn.Read(2)
	var cmd Command
	if err != nil {
		return nil, make([]byte, 0), err
	}
	if len(data) > 2 {
		return nil, make([]byte, 0), fmt.Errorf("length data is < 2")
	}
	cmd = Command(data[0])
	length := calcLength(data[1])
	if length != 0 {
		data2, err := prot.conn.Read(length)
		if err != nil {
			return nil, make([]byte, 0), err
		}
		data = append(data, data2...)
	}
	return &cmd, data, nil

}

func (prot *MqttProtocol) isMqttCmd(Cmd Command) ([]byte, error) {
	data, err := prot.conn.Read(2)
	if err != nil {
		return make([]byte, 0), err
	}
	if len(data) > 2 {
		return make([]byte, 0), fmt.Errorf("length data is < 2")
	}
	if data[0]&byte(Cmd) != byte(Cmd) {
		return make([]byte, 0), fmt.Errorf("Isn't a first byte %d Mqtt protocol", Cmd)
	}
	length := calcLength(data[1])
	data2, err := prot.conn.Read(length)
	data = append(data, data2...)
	if !mqttVersionCompatible(data[2:9]) {
		return make([]byte, 0), fmt.Errorf("Invalid MQTT Protocol Name or Version")
	}
	return data, nil
}

func (prot *MqttProtocol) ConnectProcess() (*ResponseConnect, error) {
	r, err := prot.connectUnPack()
	if err != nil {
		if r == nil {
			return nil, err
		}
		err = prot.connAck(r, CONNECT_UNACCEPCTABLE_PROTOCOL)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if r == nil { //Clean Session False
		err = prot.connAck(r, CONNECT_ACCEPCTED)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	err = prot.connAck(r, CONNECT_ACCEPCTED)
	if err != nil {
		return nil, err
	}
	return r, nil
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
	fmt.Printf("Payload: %v\n", payload)
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
		log.Println("hasWillFlag")
		//lengthWill = 0
		payload = payload[lengthWill:]
	}
	if hasUser {
		log.Println("hasUser")
		lengthUser = int(payload[0])<<8 + int(payload[1]) //TODO calc length
		payload = payload[2:]
		username = string(payload[:lengthUser])
		payload = payload[lengthUser:]
	}
	if hasPassword {
		log.Println("hasPassword")
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
	response := []byte{byte(CONNACK), 0b10}
	if sessionCfg.Clean {
		response = append(response, []byte{0x00}...)
	} else {
		response = append(response, []byte{0x01}...)
	}
	response = append(response, []byte{byte(resp)}...) //TODO response
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func (prot *MqttProtocol) publishUnPack(data []byte) (*ResponsePublish, error) {
	var response ResponsePublish
	fmt.Println(data)
	response.dutFlag = (data[0]&0b1000 == 0b1000)
	response.Qos = int(data[0] & 0b110)
	response.Retained = (data[0]&0b1 == 0b1)
	variableHeader := data[2:]
	log.Println(variableHeader)
	if len(variableHeader) < 3 {
		return nil, fmt.Errorf("tamanho do payload(%d) sem topic: %v", len(variableHeader), variableHeader)
	}
	lengthTopic := int(variableHeader[0])<<8 + int(variableHeader[1])
	variableHeader = variableHeader[2:]
	response.Topic = string(variableHeader[:lengthTopic])
	log.Println(response.Topic)
	variableHeader = variableHeader[lengthTopic:]
	if response.Qos > 0 {
		response.Identifier = int(variableHeader[0])<<8 + int(variableHeader[1])
		log.Println(response.Identifier)
		variableHeader = variableHeader[2:]
	}
	payload := variableHeader
	response.Body = payload
	log.Println(string(response.Body))
	return &response, nil
}
func (prot *MqttProtocol) pubAck(pubCfg *ResponsePublish, resp PublishReturnCode) error {
	return nil
}
func (prot *MqttProtocol) pubRec(pubCfg *ResponsePublish, resp PublishReturnCode) error {
	return nil
}

func (prot *MqttProtocol) PublishProcess(data []byte) (*ResponsePublish, error) {
	r, err := prot.publishUnPack(data)
	if err != nil {
		return nil, err
	}
	if r.Qos == 1 {
		err = prot.pubAck(r, PUBLISHED)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	if r.Qos == 2 {
		err = prot.pubRec(r, PUBLISHED)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return r, nil
}

func (prot *MqttProtocol) PingProcess() error {
	response := []byte{byte(PINGRESP), 0}
	err := prot.conn.Write(response)
	if err != nil {
		return err
	}
	return nil
}
