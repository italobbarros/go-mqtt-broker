package broker

import (
	"strings"
)

// PrintTree imprime a árvore de tópicos a partir do nó fornecido
func (b *Broker) PrintAllTree() {
	//b.printTree(b.Root, 0, "")
}

// Função auxiliar para imprimir a árvore recursivamente
func (b *Broker) printTree(node *TopicNode, level int, arrow string) {
	indent := strings.Repeat("  ", level)

	if node.TopicConfig != nil {
		b.logger.Debug(
			indent+arrow+" "+node.Name+" -> "+"Data:%s | Retained: %v | Qos: %d | Subscribers: %d",
			string(node.TopicConfig.Payload), node.TopicConfig.Retained, node.TopicConfig.Qos, len(node.Subscribers),
		)
	} else {
		b.logger.Debug(indent + arrow + " " + node.Name)
	}

	/*
		for _, child := range node.Children {
			var newArrow string
			if index == len(node.Children)-1 {
			newArrow = "└──"
			} else {
			newArrow = "├──"
			}
			b.printTree(child, level+1, newArrow)
		}*/
}

func getTopicUntilKeyword(topic, keyword string) string {
	parts := strings.Split(topic, "/")
	var result []string

	for i, part := range parts {
		result = append(result, part)
		if part == keyword && i != len(parts)-1 {
			break
		}
	}

	return strings.Join(result, "/")
}
