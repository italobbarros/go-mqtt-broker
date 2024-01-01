package utils

import (
	"flag"
	"fmt"
)

func ConfigArgs() (int, int, int) {
	timeout := flag.Int("timeoutDev", 3600, "an int")
	portMqtt := flag.Int("portMqtt", 1883, "an int")
	portManagement := flag.Int("portManagement", 43038, "an int")
	flag.Parse()
	fmt.Println("portManagement:", *portManagement)
	fmt.Println("portMqtt:", *portMqtt)
	return *portManagement, *portMqtt, *timeout
}
