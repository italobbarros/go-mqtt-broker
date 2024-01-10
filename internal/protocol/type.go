package protocol

import (
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

type MqttProtocol struct {
	conn   connection.ConnectionInterface
	logger *logger.Logger
}
