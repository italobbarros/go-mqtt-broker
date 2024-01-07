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
	defer func() {
		conn.Close()
		log.Println("Closing client MQTT")

	}()
	prot := protocol.NewMqttProtocol(conn)
	sessionCfg, err := prot.ConnectProcess()
	if err != nil {
		log.Println(err)
		return
	}
	b.newSession(sessionCfg)
	for {
		cmd, data, err := prot.IsValidMqttCmd()
		if err != nil {
			log.Println(err)
			return
		}
		if cmd == nil {
			log.Println("Command is nil")
			return
		}
		if *cmd&protocol.PUBLISH == protocol.PUBLISH {
			err := b.handlePublishCommand(data, prot)
			if err != nil {
				log.Printf("handlePublishCommand error: %s\n", err)
				return
			}
			continue
		}
		if *cmd&protocol.PINGREQ == protocol.PINGREQ {
			err := prot.PingProcess()
			if err != nil {
				log.Printf("PingProcess error: %s\n", err)
				return
			}
			log.Println("PING!")
			continue
		}
	}
}

func (b *Broker) newSession(sessionCfg *protocol.ResponseConnect) {
	defer b.SessionMg.DebugPrint()
	if sessionCfg == nil {
		return
	}
	if b.SessionMg.Exist(sessionCfg.Id) {
		b.SessionMg.UpdateSession(&SessionConfig{
			Id:        sessionCfg.Id,
			KeepAlive: sessionCfg.KeepAlive,
			username:  sessionCfg.Username,
			password:  sessionCfg.Password,
		})
		return
	}
	b.SessionMg.AddSession(&SessionConfig{
		Id:        sessionCfg.Id,
		KeepAlive: sessionCfg.KeepAlive,
		username:  sessionCfg.Username,
		password:  sessionCfg.Password,
	})
	return
}

func (b *Broker) handlePublishCommand(data []byte, prot *protocol.MqttProtocol) error {
	r, err := prot.PublishProcess(data)
	if err != nil {
		return err
	}
	b.AddTopic(r.Topic, &TopicConfig{
		Retained: r.Retained,
		Payload:  string(r.Payload),
		Qos:      r.Qos,
	})
	return nil
}

func (b *Broker) Start() {
	log.Print("Starting broker...")
	for conn := range b.server.GetChannel() {
		if conn != nil {
			go b.handleConnectionMQTT(conn)
		}
	}
}
