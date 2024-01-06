package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (prot *MqttProtocol) isMqttCmd(Cmd Command) ([]byte, error) {
	data, err := prot.conn.ReadExactly(2)
	if err != nil {
		return make([]byte, 0), err
	}
	if len(data) > 2 {
		return make([]byte, 0), fmt.Errorf("length data is < 2")
	}
	if data[0] != byte(Cmd) {
		return make([]byte, 0), fmt.Errorf("Isn't a first byte connect Mqtt protocol")
	}
	length := calcLength(data[1])
	data2, err := prot.conn.ReadExactly(length)
	data = append(data, data2...)
	if mqttVersionCompatible(data[2:9]) {
		return make([]byte, 0), fmt.Errorf("Invalid MQTT Protocol Name or Version")
	}
	return data, nil
}

func (prot *MqttProtocol) connectUnPack() (ResponseConnect, error) {
	var response ResponseConnect
	data, err := prot.isMqttCmd(CONNECT)
	if err != nil {
		return response, err
	}
	msg := data[9:]
	if len(msg) < 3 {
		return response, fmt.Errorf("length data is < 3")
	}
	connectFlag := msg[0]
	if connectFlag&0b1 == 0b1 {
		return response, fmt.Errorf("Reserved flag on connectFlag isn't 0")
	}
	if connectFlag&0b10 != 0b10 { //CleanSession is 0
		return response, nil
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
		return response, fmt.Errorf("Erro ao ler bytes:%s", err)
	}
	var clientId, username, password string
	var lengthWill, lengthUser, lengthClientId, lengthPayload int

	payload := msg[3:]
	lengthPayload = len(payload)
	lengthClientId = 23
	if lengthPayload < 1 {
		//TODO criar client Id random
		//("Nao forneceu ClientID")
		clientId = "ClientId_Random"
		lengthClientId = len(clientId)
	}
	clientId = string(payload[:lengthClientId])
	payload = payload[lengthClientId+1:]
	//TODO will msg
	if hasWillFlag {
		//lengthWill = 0
		payload = payload[lengthWill+1:]
	}
	if hasUser {
		lengthUser = int(payload[0])<<8 + int(payload[1]) //TODO calc length
		username = string(payload[:lengthUser])
		payload = payload[lengthUser+1:]
	}
	if hasPassword {
		lengthPassword := int(payload[0])<<8 + int(payload[1])
		password = string(password[2:lengthPassword])
	}
	response = ResponseConnect{
		Id:        clientId,
		Timeout:   1,
		Clean:     true,
		keepAlive: keepAlive,
		username:  username,
		password:  password,
	}
	return response, nil
}

func (prot *MqttProtocol) connAck(sessionCfg *ResponseConnect) []byte {
	response := []byte{byte(CONNACK), 2}
	response = append(response, []byte{0x00, 0x1}...) //TODO response
	return response
}

//func (b *Broker) isPublish(topic string, topicCfg *TopicConfig)
