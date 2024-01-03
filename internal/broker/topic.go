package broker

import (
	"strings"
)

func (b *Broker) AddTopic(topic string) {
	segments := strings.Split(topic, "/")

	currentNode := b.Root
	for index, segment := range segments {
		found := false
		for _, child := range currentNode.Children {
			if child.Name == segment {
				currentNode = child
				found = true
				break
			}
		}

		if !found {
			topicRootWithoutDash := b.Root.Topic[:len(b.Root.Topic)-1]
			newTopic := topicRootWithoutDash + getTopicUntilKeyword(topic, segment)

			// Se for o último segmento, crie o nó com um TopicConfig
			var topicConfig *TopicConfig = nil
			if index == len(segments)-1 {
				topicConfig = &TopicConfig{
					TopicName:    newTopic,
					QoS:          1, // Exemplo de valor, ajuste conforme necessário
					Retained:     true,
					Subscribers:  []string{},  // Lista vazia ou ajuste conforme necessário
					SecurityRule: "ALLOW_ALL", // Exemplo de regra de segurança, ajuste conforme necessário
				}
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
}
