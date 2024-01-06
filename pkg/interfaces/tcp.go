package connection

import (
	"fmt"
	"net"
)

// Tcp é uma implementação da interface Communicator para TCP.
type Tcp struct {
	conn net.Conn
}

// NewTcp cria uma nova instância de Tcp.
func NewTcp(address string) (*Tcp, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Tcp{conn: conn}, nil
}

// ReadExactly lê exatamente n bytes da conexão TCP.
func (t *Tcp) Read(length int) ([]byte, error) {
	buffer := make([]byte, length)
	n, err := t.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, fmt.Errorf("didn't read exactly %d bytes", length)
	}
	return buffer, nil
}

// Write escreve dados na conexão TCP.
func (t *Tcp) Write(data []byte) error {
	_, err := t.conn.Write(data)
	return err
}

// Close fecha a conexão TCP.
func (t *Tcp) Close() error {
	return t.conn.Close()
}

func main() {
	// Exemplo de uso:
	tcpComm, err := NewTcp("localhost:8080")
	if err != nil {
		panic(err)
	}
	defer tcpComm.Close()

	// Agora você pode usar tcpComm para ler, escrever e fechar a conexão.
}
