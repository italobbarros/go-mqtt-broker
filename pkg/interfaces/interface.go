package connection

type Communicator interface {
	Read(int) ([]byte, error)
	Write([]byte) error
	Close() error
}

type TypeConnection int

const (
	TCP TypeConnection = iota
	WEBSOCKET
)
