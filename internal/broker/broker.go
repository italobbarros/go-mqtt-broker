package broker

import (
	"fmt"

	"github.com/italobbarros/go-mqtt-broker/internal/api/models"
	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

// NewBroker inicializa um novo corretor MQTT com um nó raiz
func NewBroker(b *BrokerConfigs) *Broker {
	//topic := fmt.Sprintf("/%s/#", b.Name)
	sessionMg := NewSessionManager()
	var server connection.ServerInterface

	switch b.TypeConnection {
	case connection.TCP:
		server = connection.NewTcpServer()
	case connection.WEBSOCKET:
		//server = connection.NewTcpServer()
	}
	broker := &Broker{
		SessionMg: sessionMg,
		server:    server,
		logger:    logger.NewLogger("Broker"),
	}
	go server.Start(b.Address, broker.handleConnectionMQTT)
	return broker
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
		b.logger.Error("Error connectProcess : %s", err.Error())
		return
	}
	chSession := make(chan *Session)
	go b.newSession(sessionCfg, chSession)
	currentSession := <-chSession
	if currentSession == nil {
		return
	}
	currentSession.prot = prot
	prot.UpdateLogger(currentSession.logger)
	conn.UpdateLogger(currentSession.logger)
	b.logger.Info("Client MQTT Connected!")
	topicReady := make(chan bool)
	for {
		r := prot.IsValidMqttCmd()
		err := r.Err
		cmd := r.Command
		data := r.Data
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
			responsePublish, err = b.handlePublishCommand(data, prot, topicReady)
			if err != nil {
				prot.End()
				currentSession.logger.Error("handlePublishCommand: %s", err.Error())
				return
			}
			if responsePublish == nil {
				prot.End()
				currentSession.logger.Error("handlePublishCommand: responsePublish == nil")
				return
			}
			if responsePublish.Qos < 2 {
				prot.End()
			}
			go b.AddPublish(models.PublishRequest{
				ClientIdSession: currentSession.Id,
				Payload:         string(responsePublish.Payload),
				Qos:             responsePublish.Qos,
				TopicName:       responsePublish.Topic,
				TopicRetained:   responsePublish.Retained,
			})

			continue
		}
		if protocol.IsCmdEqual(cmd, protocol.COMMAND_PUBREL) && responsePublish != nil { // exactly equal
			//Continued command publish if qos is 2
			iden := responsePublish.Identifier
			err := prot.PubRelProcess(data, &iden)
			if err != nil {
				prot.End()
				currentSession.logger.Error("PubRelProcess: %s", err.Error())
				return
			}
			go b.NotifyAllSubscribers(responsePublish.Topic, topicReady)
			responsePublish = nil
			currentSession.logger.Info("Published!")
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
				b.logger.Debug("AddSubscribeTopicNode: %s", topic)
				err = b.AddSubscribeTopicNode(
					topic,
					currentSession.Id,
					&SubscriberConfig{
						Identifier: subs.Identifier,
						Qos:        subs.Qos[index],
						session:    currentSession,
					}, topicReady,
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
		if protocol.IsCmdEqual(cmd, protocol.COMMAND_UNSUBSCRIBE) {
			prot.Start()
			subs, err := prot.UnSubscribeProcess(data)
			if err != nil {
				prot.End()
				currentSession.logger.Error("SubscribeProcess: %s", err.Error())
				return
			}
			for _, topic := range subs.TopicFilter {
				b.logger.Debug("remove subscribe: %s", topic)
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

func (b *Broker) newSession(sessionCfg *protocol.ResponseConnect, chSession chan *Session) {
	//defer b.SessionMg.DebugPrint()
	if sessionCfg == nil {
		chSession <- nil
		return
	}
	b.SessionMg.AddSession(&SessionConfig{
		Id:        sessionCfg.Id,
		KeepAlive: sessionCfg.KeepAlive,
		username:  sessionCfg.Username,
		password:  sessionCfg.Password,
	}, chSession)
}

func (b *Broker) handlePublishCommand(data []byte, prot *protocol.MqttProtocol, topicReady chan bool) (*protocol.ResponsePublish, error) {
	r, err := prot.PublishProcess(data)
	if err != nil {
		return nil, err
	}
	if r.Qos == 2 {
		b.logger.Debug("Publish Qos 2")
		return r, nil
	}
	//go b.NotifyAllSubscribers(r.Topic, topicReady)
	return r, nil
}

func (b *Broker) Start() {
	fmt.Print("Starting broker...")
	for conn := range b.server.GetChannel() {
		if conn != nil {
			go b.handleConnectionMQTT(conn)
		}
	}
}
