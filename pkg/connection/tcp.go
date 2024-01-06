package connection

import (
	"fmt"
	"log"
	"net"
	"time"
)

type TcpConn struct {
	conn net.Conn
}

type TcpServer struct {
	ChNewConn chan ConnectionInterface
}

func NewTcpConn(conn net.Conn) *TcpConn {
	tcp := TcpConn{conn: conn}
	return &tcp
}

// NewTcpServer cria uma nova instância de TcpServer.
func NewTcpServer() *TcpServer {
	return &TcpServer{
		ChNewConn: make(chan ConnectionInterface),
	}
}

func (t *TcpServer) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("Erro ao fazer o listen: %s", err)
		time.Sleep(time.Second * 60)
		t.Start(address)
	}
	defer listener.Close()

	fmt.Println("Servidor TCP ouvindo ->", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		tcpConn := NewTcpConn(conn)
		t.ChNewConn <- tcpConn
	}
}

func (t *TcpServer) GetChannel() chan ConnectionInterface {
	return t.ChNewConn
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
	log.Printf("Read: %v\n", buffer)
	return buffer, nil
}

// Write escreve dados na conexão TCP.
func (t *TcpConn) Write(data []byte) error {
	log.Printf("Write: %v\n", data)
	_, err := t.conn.Write(data)
	return err
}

// Close fecha a conexão TCP.
func (t *TcpConn) Close() error {
	return t.conn.Close()
}
