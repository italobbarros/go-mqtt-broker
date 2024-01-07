package main

import (
	"github.com/italobbarros/go-mqtt-broker/internal/api"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

func main() {
	logger.InitCustomFormatter()
	broker := brokerMqtt.NewBroker(&brokerMqtt.BrokerConfigs{
		Name:           "go-mqtt-broker",
		Address:        "0.0.0.0:1883",
		TypeConnection: connection.TCP,
	})
	go broker.Start()
	api := api.NewAPI(broker)
	go api.Init()
	select {}
}
