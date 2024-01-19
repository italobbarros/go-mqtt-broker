package main

import (
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/italobbarros/go-mqtt-broker/internal/api"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
	"github.com/italobbarros/go-mqtt-broker/internal/utils"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

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

	select {
	case sig := <-sigChan:
		handleSignal(sig, broker)
	}
}
func handleSignal(sig os.Signal, broker *brokerMqtt.Broker) {
	var wg sync.WaitGroup
	switch sig {
	case syscall.SIGINT:
		broker.DisconnectAllSessions(&wg)
	case syscall.SIGTERM:
		broker.DisconnectAllSessions(&wg)
	}
	wg.Wait()
	os.Exit(0)
}
