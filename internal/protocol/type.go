package protocol

import interfaces "github.com/italobbarros/go-mqtt-broker/pkg/connection"

type Command int

const (
	CONNECT Command = 0x10
	CONNACK Command = 0x20
	Publish
	UnPublish
	Subscribe
	UnSubscribe
)

type MqttProtocol struct {
	conn interfaces.Communicator
}

func NewMqttProtocol(conn interfaces.Communicator) *MqttProtocol {
	return &MqttProtocol{
		conn: conn,
	}
}
