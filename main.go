package main

import (
	"fmt"

	"github.com/italobbarros/go-mqtt-broker/internal/api"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
)

func main() {
	broker := brokerMqtt.NewBroker("Container1")
	api := api.NewAPI(broker)
	go api.Init()
	topics := []string{
		"client1/teste/io1",
		"client1/teste/io2",
		"client2/val/io1",
		"client2/val/io2",
		"client3/io1",
		"client3/io2",
	}

	for _, topic := range topics {
		fmt.Println("add: " + topic)
		broker.AddTopic(topic)
	}
	// Exibir a árvore (apenas para demonstração)
	broker.PrintAllTree()
	select {}
}
