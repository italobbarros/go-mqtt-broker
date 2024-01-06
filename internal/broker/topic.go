package broker

import (
	"fmt"
	"strings"
)

func (b *Broker) AddTopic(topic string, topicCfg *TopicConfig) {
	fmt.Println("add: " + topic)
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
