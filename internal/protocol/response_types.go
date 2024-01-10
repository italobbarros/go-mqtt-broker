package protocol

import "github.com/italobbarros/go-mqtt-broker/pkg/connection"

type ResponseConnect struct {
	Id        string
	Timeout   int
	KeepAlive int16
	Clean     bool
	Username  string
	Password  string
}

type ResponsePublish struct {
	Identifier []byte
	Topic      string
	Payload    []byte
	dutFlag    bool
	Retained   bool
	Qos        int
}

type ResponseSubscribe struct {
	Identifier  []byte
	TopicFilter []string
	Qos         []int
	Conn        connection.ConnectionInterface
}
