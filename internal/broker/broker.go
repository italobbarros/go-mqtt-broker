package broker

import (
	"fmt"
	"log"
	"time"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	interfaces "github.com/italobbarros/go-mqtt-broker/pkg/interfaces"
)

// NewBroker inicializa um novo corretor MQTT com um nó raiz
func NewBroker(b *BrokerConfigs) *Broker {
	topic := fmt.Sprintf("/%s/#", b.Name)
	sessionMg := NewSessionManager()
	var conn interfaces.Communicator
	var err error
	switch b.TypeConnection {
	case interfaces.TCP:
		conn, err = interfaces.NewTcp(b.Address)
		if err != nil {
			log.Printf("Erro ao se conectar na tcp%s", err)
			time.Sleep(time.Second * 60)
			return NewBroker(b)
		}
	default:
		conn, err = interfaces.NewTcp(b.Address)
		if err != nil {
			log.Printf("Erro ao se conectar na tcp%s", err)
			time.Sleep(time.Second * 60)
			return NewBroker(b)
		}
	}
	prot := protocol.NewMqttProtocol(conn)
	return &Broker{
		Root: &TopicNode{
			Name:     b.Name,
			Topic:    topic,
			Children: make([]*TopicNode, 0),
		},
		SessionMg: sessionMg,
		protocol:  prot,
	}
}
