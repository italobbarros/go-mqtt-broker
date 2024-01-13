package broker

import (
	"fmt"
	"strings"
	"time"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
)

func (b *Broker) AddTopic(topic string, topicCfg *TopicConfig, topicReady chan bool) {
	segments := strings.Split(topic, "/")

	currentNode := b.Root
	for index, segment := range segments {
		currentNode.lockChildren.RLock()
		child, ok := currentNode.Children[segment]
		currentNode.lockChildren.RUnlock()

		if ok {
			if child.Name == segment {
				// Se for o último segmento, atualize o TopicConfig existente
				if index == len(segments)-1 {
					child.TopicConfig = topicCfg
				}
				currentNode = child
				continue
			}
		}
		newTopic := getTopicUntilKeyword(topic, segment)
		// Se for o último segmento, crie o nó com um TopicConfig
		var topicConfig *TopicConfig = nil
		if index == len(segments)-1 {
			topicConfig = topicCfg
		}
		newChild := &TopicNode{
			Name:        segment,
			Topic:       newTopic,
			TopicConfig: topicConfig,
			Children:    make(map[string]*TopicNode),
			Subscribers: make(map[string]*SubscriberConfig),
		}
		currentNode.lockChildren.Lock()
		currentNode.Children[segment] = newChild
		currentNode.lockChildren.Unlock()
		currentNode = newChild
	}
	topicReady <- true
	//b.PrintAllTree()
}

func (b *Broker) GetTopicNode(topic string) *TopicNode {
	segments := strings.Split(topic, "/")
	currentNode := b.Root
	for _, segment := range segments {
		currentNode.lockChildren.RLock()
		child, ok := currentNode.Children[segment]
		currentNode.lockChildren.RUnlock()
		if !ok {
			return nil
		}
		currentNode = child
	}
	return currentNode
}

func (b *Broker) AddSubscribeTopicNode(topic string, id string, subs *SubscriberConfig, topicNode chan bool) error {
	TopicNode := b.GetTopicNode(topic)
	if TopicNode == nil {
		b.logger.Warning("Don't exist Topic on Topic Node")
		b.AddTopic(topic, &TopicConfig{
			Retained: false,
			Payload:  "",
			Qos:      0,
		}, topicNode)
		TopicNode = b.GetTopicNode(topic)
	}
	TopicNode.Subscribers[id] = subs
	TopicNode.SubscribersCount++
	b.logger.Debug("Add subscriber...")
	b.logger.Debug("SubscribersCount: %d", TopicNode.SubscribersCount)
	if TopicNode.TopicConfig.Retained {
		if err := b.notifyNewSubscriber(topic, subs); err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) notifyNewSubscriber(topic string, sub *SubscriberConfig) error {
	b.logger.Debug("NotifyNewSubscriber: %s", sub.Identifier)
	return nil
}

func (b *Broker) NotifyAllSubscribers(topic string, topicReady chan bool) {
	<-topicReady
	var topicNode *TopicNode
	for {
		topicNode = b.GetTopicNode(topic)
		if topicNode != nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
	topicNode.MessageCount += 1
	for _, sub := range topicNode.Subscribers {
		currentProt := sub.session.prot
		go b.publishSubscribe(topicNode.Topic, currentProt, topicNode.TopicConfig.Retained, topicNode.TopicConfig.Payload, sub)
	}
}

func (b *Broker) publishSubscribe(topic string, currentProt *protocol.MqttProtocol, retained bool, payload string, sub *SubscriberConfig) {
	currentProt.Start()
	err := currentProt.PublishSend(sub.Qos, true, retained, payload, topic, sub.Identifier)
	if err != nil {
		currentProt.End()
	}
	b.logger.Debug("sub.Identifier: %s", sub.Identifier)
	currentProt.End()
}

func (b *Broker) IncMsgCount(topic string) error {
	TopicNode := b.GetTopicNode(topic)
	if TopicNode == nil {
		return fmt.Errorf("Don't exist Topic on Topic Node")
	}
	TopicNode.MessageCount += 1
	return nil
}
