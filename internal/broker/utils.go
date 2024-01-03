package broker

import (
	"fmt"
	"strings"
)

// PrintTree imprime a árvore de tópicos a partir do nó fornecido
func (b *Broker) PrintAllTree() {
	printTree(b.Root, 0)
}

// Função auxiliar para imprimir a árvore recursivamente
func printTree(node *TopicNode, level int) {
	indent := strings.Repeat("  ", level)
	fmt.Println(indent + node.Name)
	for _, child := range node.Children {
		printTree(child, level+1)
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
