package connection

import (
	"fmt"
	"net"
	"time"

	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

type TcpConn struct {
	conn   net.Conn
	logger *logger.Logger
}

type TcpServer struct {
	ChNewConn chan ConnectionInterface
	logger    *logger.Logger
}

func NewTcpConn(conn net.Conn) *TcpConn {
	tcp := TcpConn{
		conn:   conn,
		logger: logger.NewLogger("TCP"),
	}
	return &tcp
}

// NewTcpServer cria uma nova instância de TcpServer.
func NewTcpServer() *TcpServer {
	return &TcpServer{
		ChNewConn: make(chan ConnectionInterface),
		logger:    logger.NewLogger("TCPServer"),
	}
}

func (t *TcpServer) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.logger.Error("Erro ao fazer o listen: %s", err)
		time.Sleep(time.Second * 60)
		t.Start(address)
	}
	defer listener.Close()

	t.logger.Debug("Servidor TCP ouvindo ->", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.logger.Debug("Erro ao aceitar conexão:", err)
			continue
		}
		tcpConn := NewTcpConn(conn)
		t.ChNewConn <- tcpConn
	}
}

func (t *TcpServer) GetChannel() chan ConnectionInterface {
	return t.ChNewConn
}
func (t *TcpConn) UpdateLogger(logger *logger.Logger) {
	t.logger = logger
}

// ReadExactly lê exatamente n bytes da conexão TCP.
func (t *TcpConn) Read(length int) ([]byte, error) {
	buffer := make([]byte, length)
	n, err := t.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, fmt.Errorf("didn't read exactly %d bytes", length)
	}
	t.logger.Answer("Read: %v", buffer)
	return buffer, nil
}

// Write escreve dados na conexão TCP.
func (t *TcpConn) Write(data []byte) error {
	err := t.conn.SetWriteDeadline(time.Now().Add(time.Duration(30) * time.Second))
	if err != nil {
		return err
	}
	t.logger.Ask("Write: %v", data)
	_, err = t.conn.Write(data)
	return err
}

// Close fecha a conexão TCP.
func (t *TcpConn) Close() error {
	t.logger.Debug("Closing...")
	return t.conn.Close()
}
