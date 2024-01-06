package broker

import (
	"fmt"
	"strings"
)

// PrintTree imprime a árvore de tópicos a partir do nó fornecido
func (b *Broker) PrintAllTree() {
	printTree(b.Root, 0, "")
}

// Função auxiliar para imprimir a árvore recursivamente
func printTree(node *TopicNode, level int, arrow string) {
	indent := strings.Repeat("  ", level)

	if node.TopicConfig != nil {
		fmt.Print(indent + arrow + " " + node.Name)
		fmt.Print(" -> "+"Data:", string(node.TopicConfig.Data))
		fmt.Print(" | "+"Retained:", node.TopicConfig.Retained)
		fmt.Print(" | "+"Subscribers:", node.TopicConfig.Subscribers)
		fmt.Println(" | "+"SecurityRule:", node.TopicConfig.SecurityRule)
	} else {
		fmt.Println(indent + arrow + " " + node.Name)
	}

	for index, child := range node.Children {
		var newArrow string
		if index == len(node.Children)-1 {
			newArrow = "└──"
		} else {
			newArrow = "├──"
		}
		printTree(child, level+1, newArrow)
	}
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
