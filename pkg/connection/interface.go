package connection

type Communicator interface {
	Read() ([]byte, error)
	ReadExactly(int) ([]byte, error)
	Write([]byte) error
	Close() error
}
