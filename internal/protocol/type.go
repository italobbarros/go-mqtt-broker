package protocol

import interfaces "github.com/italobbarros/go-mqtt-broker/pkg/interfaces"

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
