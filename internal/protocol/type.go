package protocol

import connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"

type Command int

const (
	CONNECT Command = 0x10
	CONNACK Command = 0x20
	Publish
	UnPublish
	Subscribe
	UnSubscribe
)

type ConnectReturnCode byte

const (
	ACCEPCTED              ConnectReturnCode = 0x00 //Connection Accepted Connection accepted
	UNACCEPCTABLE_PROTOCOL ConnectReturnCode = 0x01 //Connection Refused, unacceptable protocol version The Server does not support the level of the MQTT protocol requested by the Client
	REJECTED_CLIENT        ConnectReturnCode = 0x02 //Connection Refused, identifier rejected The Client identifier is correct UTF-8 but not allowed by the Server
	SERVER_UNAVAILABLE     ConnectReturnCode = 0x03 //Connection Refused, Server unavailable The Network Connection has been made but the MQTT service is unavailable
	USER_OR_PASS_MALFORMED ConnectReturnCode = 0x04 //Connection Refused, bad user name or password The data in the user name or password is malformed
	NOT_AUTHORIZED         ConnectReturnCode = 0x05 //Connection Refused, not authorized The Client is not authorized to connect
)

type MqttProtocol struct {
	conn connection.ConnectionInterface
}
