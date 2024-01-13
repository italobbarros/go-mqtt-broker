package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/italobbarros/go-mqtt-broker/internal/api"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

func main() {
	logger.InitCustomFormatter()
	broker := brokerMqtt.NewBroker(&brokerMqtt.BrokerConfigs{
		Name:           "go-mqtt-broker",
		Address:        os.Getenv("ADDRESS"),
		TypeConnection: connection.TCP,
	})
	go broker.Start()
	api := api.NewAPI(broker)
	go api.Init()
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	select {}
}
