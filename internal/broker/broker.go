package broker

import (
	"fmt"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
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
		logger:    logger.NewLogger("Broker"),
	}
}

func (b *Broker) handleConnectionMQTT(conn connection.ConnectionInterface) {
	var packetIdentifier *[]byte = nil
	b.logger.Debug("Client Connecting...")
	defer func() {
		conn.Close()
		b.logger.Debug("Closing client MQTT")
	}()
	prot := protocol.NewMqttProtocol(conn)
	sessionCfg, err := prot.ConnectProcess()
	if err != nil {
		b.logger.Error(err.Error())
		return
	}
	currentSession := b.newSession(sessionCfg)
	prot.UpdateLogger(currentSession.logger)
	conn.UpdateLogger(currentSession.logger)
	for {
		cmd, data, err := prot.IsValidMqttCmd()
		if err != nil {
			currentSession.logger.Error(err.Error())
			return
		}
		if cmd == nil {
			currentSession.logger.Error("Command is nil")
			return
		}
		if *cmd&protocol.COMMAND_PUBLISH == protocol.COMMAND_PUBLISH {
			packetIdentifier, err = b.handlePublishCommand(data, prot)
			if err != nil {
				currentSession.logger.Error("handlePublishCommand: %s", err.Error())
				return
			}
			currentSession.logger.Debug("Published!")
			continue
		}
		if (*cmd&protocol.COMMAND_PUBREL == protocol.COMMAND_PUBREL) && (packetIdentifier != nil) { // exactly equal
			err := prot.PubRelProcess(data, packetIdentifier)
			if err != nil {
				currentSession.logger.Error("PubRelProcess: %s", err.Error())
				return
			}
			packetIdentifier = nil
			currentSession.logger.Debug("Success PubRel!")
			continue
		}
		if *cmd&protocol.COMMAND_PINGREQ == protocol.COMMAND_PINGREQ {
			err := prot.PingProcess()
			if err != nil {
				currentSession.logger.Error("PingProcess: %s", err.Error())
				return
			}
			currentSession.logger.Debug("PING!")
			continue
		}
	}
}

func (b *Broker) newSession(sessionCfg *protocol.ResponseConnect) *Session {
	var currentSession *Session = nil

	defer b.SessionMg.DebugPrint()
	if sessionCfg == nil {
		return currentSession
	}
	if b.SessionMg.Exist(sessionCfg.Id) {
		currentSession = b.SessionMg.UpdateSession(&SessionConfig{
			Id:        sessionCfg.Id,
			KeepAlive: sessionCfg.KeepAlive,
			username:  sessionCfg.Username,
			password:  sessionCfg.Password,
		})
		return currentSession
	}
	currentSession = b.SessionMg.AddSession(&SessionConfig{
		Id:        sessionCfg.Id,
		KeepAlive: sessionCfg.KeepAlive,
		username:  sessionCfg.Username,
		password:  sessionCfg.Password,
	})
	return currentSession
}

func (b *Broker) handlePublishCommand(data []byte, prot *protocol.MqttProtocol) (*[]byte, error) {
	r, err := prot.PublishProcess(data)
	if err != nil {
		return nil, nil
	}
	b.AddTopic(r.Topic, &TopicConfig{
		Retained: r.Retained,
		Payload:  string(r.Payload),
		Qos:      r.Qos,
	})
	if r.Qos == 2 {
		return &r.Identifier, nil
	}
	return nil, nil
}

func (b *Broker) Start() {
	fmt.Print("Starting broker...")
	for conn := range b.server.GetChannel() {
		if conn != nil {
			go b.handleConnectionMQTT(conn)
		}
	}
}
