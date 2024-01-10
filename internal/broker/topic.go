package broker

import (
	"fmt"
	"strings"
)

func (b *Broker) AddTopic(topic string, topicCfg *TopicConfig) {
	segments := strings.Split(topic, "/")

	currentNode := b.Root
	for index, segment := range segments {
		found := false
		for _, child := range currentNode.Children {
			if child.Name == segment {
				// Se for o último segmento, atualize o TopicConfig existente
				if index == len(segments)-1 {
					child.TopicConfig = topicCfg
				}
				currentNode = child
				found = true
				break
			}
		}

		if !found {
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
				Children:    make([]*TopicNode, 0),
			}
			currentNode.Children = append(currentNode.Children, newChild)
			currentNode = newChild
		}
	}
	b.PrintAllTree()
}

func (b *Broker) GetTopicNode(topic string) *TopicNode {
	segments := strings.Split(topic, "/")
	currentNode := b.Root

	for _, segment := range segments {
		found := false
		for _, child := range currentNode.Children {
			if child.Name == segment {
				currentNode = child
				found = true
				break
			}
		}

		// Se não encontrou o segmento atual, retorna nil
		if !found {
			return nil
		}
	}

	return currentNode
}

func (b *Broker) AddSubscribeTopicNode(topic string, subs *SubscriberConfig) error {
	TopicNode := b.GetTopicNode(topic)
	if TopicNode == nil {
		return fmt.Errorf("Don't exist Topic on Topic Node")
	}
	b.logger.Debug("Add subscriber.")
	TopicNode.Subscribers = append(TopicNode.Subscribers, []SubscriberConfig{*subs}...)
	if TopicNode.TopicConfig.Retained {
		if err := b.NotifyNewSubscriber(topic, subs); err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) NotifyNewSubscriber(topic string, subs *SubscriberConfig) error {

	return nil
}

func (b *Broker) NotifyAllSubscribers(topic string) error {

	return nil
}
