package broker

import (
	"fmt"
	"strings"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
)

func (b *Broker) AddTopic(topic string, topicCfg *TopicConfig) {
	segments := strings.Split(topic, "/")

	currentNode := b.Root
	for index, segment := range segments {
		currentNode.lock.RLock()
		child, ok := currentNode.Children[segment]
		currentNode.lock.RUnlock()
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
		currentNode.lock.Lock()
		currentNode.Children[segment] = newChild
		currentNode.lock.Unlock()
		currentNode = newChild

	}
	//b.PrintAllTree()
}

func (b *Broker) GetTopicNode(topic string) *TopicNode {
	segments := strings.Split(topic, "/")
	currentNode := b.Root
	for _, segment := range segments {
		child, ok := currentNode.Children[segment]
		if !ok {
			return nil
		}
		currentNode = child
	}
	return currentNode
}

func (b *Broker) AddSubscribeTopicNode(topic string, id string, subs *SubscriberConfig) error {
	TopicNode := b.GetTopicNode(topic)
	if TopicNode == nil {
		b.logger.Warning("Don't exist Topic on Topic Node")
		b.AddTopic(topic, &TopicConfig{
			Retained: false,
			Payload:  "",
			Qos:      0,
		})
		TopicNode = b.GetTopicNode(topic)
	}
	TopicNode.Subscribers[id] = subs
	TopicNode.SubscribersCount = len(TopicNode.Subscribers)
	b.logger.Debug("Add subscriber...")
	b.logger.Debug("SubscribersCount: %d", len(TopicNode.Subscribers))
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

func (b *Broker) NotifyAllSubscribers(topic string) error {
	TopicNode := b.GetTopicNode(topic)
	if TopicNode == nil {
		return fmt.Errorf("Don't exist Topic on Topic Node")
	}
	for _, sub := range TopicNode.Subscribers {
		currentProt := sub.session.prot
		go b.publishSubscribe(topic, currentProt, TopicNode.TopicConfig.Retained, TopicNode.TopicConfig.Payload, sub)
	}
	return nil
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
