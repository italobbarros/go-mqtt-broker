package connection

import "github.com/italobbarros/go-mqtt-broker/pkg/logger"

type ConnectionInterface interface {
	Read(int) ([]byte, error)
	Write([]byte) error
	Close() error
	UpdateLogger(*logger.Logger)
}

type ServerInterface interface {
	Start(string)
	GetChannel() chan ConnectionInterface
}

type TypeConnection int

const (
	TCP TypeConnection = iota
	WEBSOCKET
)
