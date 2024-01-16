package main

import (
	_ "net/http/pprof"
	"os"

	"github.com/italobbarros/go-mqtt-broker/internal/api"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
	"github.com/italobbarros/go-mqtt-broker/internal/utils"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

func main() {
	utils.LoadEnv()
	logger.InitCustomFormatter()
	broker := brokerMqtt.NewBroker(&brokerMqtt.BrokerConfigs{
		Name:           "go-mqtt-broker",
		Address:        os.Getenv("ADDRESS"),
		TypeConnection: connection.TCP,
	})
	go broker.Start()
	api := api.NewAPI()
	go api.Init()

	select {}
}
