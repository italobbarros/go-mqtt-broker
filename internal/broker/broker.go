package broker

import "fmt"

// NewBroker inicializa um novo corretor MQTT com um nó raiz
func NewBroker(name string) *Broker {
	topic := fmt.Sprintf("/%s/#", name)
	return &Broker{
		Root: &TopicNode{
			Name:     name,
			Topic:    topic,
			Children: make([]*TopicNode, 0),
		},
	}
}
