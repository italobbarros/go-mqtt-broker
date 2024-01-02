package broker

import (
	"fmt"
	"strings"
)

// AddTopic adiciona um tópico à árvore
func (b *Broker) AddTopic(topic string) {
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

		if !found {
			fmt.Print(topic + "|" + segment)
			topicWithoutWildcard := b.Root.Topic[:len(b.Root.Topic)-1]
			newChild := &TreeNode{
				Name:     segment,
				Topic:    topicWithoutWildcard + getTopicUntilKeyword(topic, segment),
				Children: make([]*TreeNode, 0),
			}
			currentNode.Children = append(currentNode.Children, newChild)
			currentNode = newChild
		}
	}
}
