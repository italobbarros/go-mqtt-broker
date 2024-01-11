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
	var responsePublish *protocol.ResponsePublish = nil
	b.logger.Debug("Client Connecting...")
	defer func() {
		conn.Close()
		b.logger.Warning("Closing client MQTT")
	}()
	prot := protocol.NewMqttProtocol(conn)
	sessionCfg, err := prot.ConnectProcess()
	if err != nil {
		b.logger.Error(err.Error())
		return
	}
	currentSession := b.newSession(sessionCfg)
	currentSession.prot = prot
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
		if protocol.IsCmdEqual(cmd, protocol.COMMAND_PUBLISH) {
			prot.Start()
			responsePublish, err = b.handlePublishCommand(data, prot)
			if err != nil {
				prot.End()
				currentSession.logger.Error("handlePublishCommand: %s", err.Error())
				return
			}
			if responsePublish == nil {
				prot.End()
			}
			currentSession.logger.Info("Published!")
			continue
		}
		if (protocol.IsCmdEqual(cmd, protocol.COMMAND_PUBREL)) && (responsePublish != nil) { // exactly equal
			//Continued command publish if qos is 2
			err := prot.PubRelProcess(data, &responsePublish.Identifier)
			if err != nil {
				prot.End()
				currentSession.logger.Error("PubRelProcess: %s", err.Error())
				return
			}
			if err = b.NotifyAllSubscribers(responsePublish.Topic); err != nil {
				b.logger.Error("Erro ao notificar todos os subscribers topic %s , erro:%s", responsePublish.Topic, err)
			}
			b.IncMsgCount(responsePublish.Topic)
			responsePublish = nil
			currentSession.logger.Info("Success PubRel!")
			prot.End()
			continue
		}
		if protocol.IsCmdEqual(cmd, protocol.COMMAND_SUBSCRIBE) {
			prot.Start()
			var Success []bool
			subs, err := prot.SubscribeProcess(data)
			if err != nil {
				prot.End()
				currentSession.logger.Error("SubscribeProcess: %s", err.Error())
				return
			}
			Success = make([]bool, len(subs.TopicFilter))
			for index, topic := range subs.TopicFilter {
				b.logger.Debug("topic: %s", topic)
				err = b.AddSubscribeTopicNode(
					topic,
					currentSession.config.Id,
					&SubscriberConfig{
						Identifier: subs.Identifier,
						Qos:        subs.Qos[index],
						session:    currentSession,
					},
				)
				if err != nil {
					b.logger.Error("Erro ao adicionar subscriber topic: %s , erro:%s", topic, err)
					Success[index] = false
				} else {
					Success[index] = true
				}
			}
			currentSession.logger.Info("Subscribed!")
			prot.End()
			continue
		}
		if protocol.IsCmdEqual(cmd, protocol.COMMAND_PINGREQ) {
			prot.Start()
			err := prot.PingProcess()
			if err != nil {
				prot.End()
				currentSession.logger.Error("PingProcess: %s", err.Error())
				return
			}
			currentSession.logger.Info("PING!")
			prot.End()
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

func (b *Broker) handlePublishCommand(data []byte, prot *protocol.MqttProtocol) (*protocol.ResponsePublish, error) {
	r, err := prot.PublishProcess(data)
	if err != nil {
		return nil, err
	}
	b.AddTopic(r.Topic, &TopicConfig{
		Retained: r.Retained,
		Payload:  string(r.Payload),
		Qos:      r.Qos,
	})
	if r.Qos == 2 {
		return r, nil
	}
	if err = b.NotifyAllSubscribers(r.Topic); err != nil {
		b.logger.Error("Erro ao notificar todos os subscribers topic %s , erro:%s", r.Topic, err)
	}
	b.IncMsgCount(r.Topic)
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
