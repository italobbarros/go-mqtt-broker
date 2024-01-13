package protocol

type Command int

const (
	CONNECT             Command = 0x10
	COMMAND_CONNACK     Command = 0x20
	COMMAND_PUBLISH     Command = 0x30
	COMMAND_PUBACK      Command = 0x40
	COMMAND_PUBREC      Command = 0x50
	COMMAND_PUBREL      Command = 0x62 // exactly equal
	COMMAND_PUBCOMP     Command = 0x70
	COMMAND_SUBSCRIBE   Command = 0x82 // exactly equal
	COMMAND_SUBACK      Command = 0x90
	COMMAND_UNSUBSCRIBE Command = 0xA2 // exactly equal
	COMMAND_UNSUBACK    Command = 0xB0
	COMMAND_PINGREQ     Command = 0xC0
	COMMAND_PINGRESP    Command = 0xD0
)

type ConnectReturnCode byte

const (
	CONNECT_ACCEPCTED              ConnectReturnCode = 0x00 //Connection Accepted Connection accepted
	CONNECT_UNACCEPCTABLE_PROTOCOL ConnectReturnCode = 0x01 //Connection Refused, unacceptable protocol version The Server does not support the level of the MQTT protocol requested by the Client
	CONNECT_REJECTED_CLIENT        ConnectReturnCode = 0x02 //Connection Refused, identifier rejected The Client identifier is correct UTF-8 but not allowed by the Server
	CONNECT_SERVER_UNAVAILABLE     ConnectReturnCode = 0x03 //Connection Refused, Server unavailable The Network Connection has been made but the MQTT service is unavailable
	CONNECT_USER_OR_PASS_MALFORMED ConnectReturnCode = 0x04 //Connection Refused, bad user name or password The data in the user name or password is malformed
	CONNECT_NOT_AUTHORIZED         ConnectReturnCode = 0x05 //Connection Refused, not authorized The Client is not authorized to connect
)

type SubscribeReturnCode byte

const (
	SUBSCRIBE_SUCCESS_QOS0 SubscribeReturnCode = 0x00 //- Success - Maximum QoS 0
	SUBSCRIBE_SUCCESS_QOS1 SubscribeReturnCode = 0x01 //- Success - Maximum QoS 1
	SUBSCRIBE_SUCCESS_QOS2 SubscribeReturnCode = 0x02 //- Success - Maximum QoS 2
	SUBSCRIBE_FAILED       SubscribeReturnCode = 0x80 //- Failure
)
