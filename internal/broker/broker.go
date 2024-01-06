package broker

import (
	"fmt"
	"log"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
)

// NewBroker inicializa um novo corretor MQTT com um nó raiz
func NewBroker(b *BrokerConfigs) *Broker {
	topic := fmt.Sprintf("/%s/#", b.Name)
	sessionMg := NewSessionManager()
	var server connection.ServerInterface
	switch b.TypeConnection {
	case connection.TCP:
		server = connection.NewTcpServer()
		go server.Start(b.Address)
	default:
		server = connection.NewTcpServer()
		go server.Start(b.Address)
	}
	return &Broker{
		Root: &TopicNode{
			Name:     b.Name,
			Topic:    topic,
			Children: make([]*TopicNode, 0),
		},
		SessionMg: sessionMg,
		server:    server,
	}
}

func (b *Broker) handleConnectionMQTT(conn connection.ConnectionInterface) {
	log.Println("handleConnectionMQTT")
	defer conn.Close()
	prot := protocol.NewMqttProtocol(conn)
	sessionCfg, err := prot.ConnectProcess()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(sessionCfg)
	for {
	}
}

func (b *Broker) Start() {
	log.Print("Starting broker...")
	for conn := range b.server.GetChannel() {
		if conn != nil {
			go b.handleConnectionMQTT(conn)
		}
	}
}
