package connection

type ConnectionInterface interface {
	Read(int) ([]byte, error)
	Write([]byte) error
	Close() error
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
