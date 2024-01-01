package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/italobbarros/go-mqtt-broker/pkg/logger"

	"github.com/italobbarros/go-mqtt-broker/internal/utils"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	utils.LoadEnv()
	logger.InitCustomFormatter()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/path1", path1Handler)
	http.HandleFunc("/path2", path2Handler)
	http.HandleFunc("/path3", path3Handler)

	fmt.Println("Servidor iniciado na porta 8080...")
	http.ListenAndServe(":8080", nil)

	//portManagement, portMqtt, timeout := utils.ConfigArgs()

	//managerDevices := device.NewManagerDevices(rdb, Db, statusSockets, timeout)
	//go managerDevices.Start(id_app, port)

	// Aguarde indefinidamente
	//select {
	//case sig := <-sigChan:
	//	handleSignal(sig, managerDevices)
	//}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bem-vindo à página raiz!</h1>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, `<li><a href="/path1">Path 1</a></li>`)
	fmt.Fprintf(w, `<li><a href="/path2">Path 2</a></li>`)
	fmt.Fprintf(w, `<li><a href="/path3">Path 3</a></li>`)
	fmt.Fprintf(w, "</ul>")
}

func path1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Você está no Path 1!</h1>")
	showSidebar(w)
}

func path2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Você está no Path 2!</h1>")
	showSidebar(w)
}

func path3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Você está no Path 3!</h1>")
	showSidebar(w)
}

func showSidebar(w http.ResponseWriter) {
	fmt.Fprintf(w, `<div style="border: 1px solid black; padding: 10px; position: fixed; right: 0; top: 0; width: 200px;">`)
	fmt.Fprintf(w, `<h3>Informações Laterais</h3>`)
	fmt.Fprintf(w, `<p>Algumas informações aqui...</p>`)
	fmt.Fprintf(w, `</div>`)
}

//func handleSignal(sig os.Signal, manDev *device.ManagerDevices) {
//	switch sig {
//	case syscall.SIGINT:
//		manDev.DisconnectAllDevices()
//		os.Exit(0)
//	case syscall.SIGTERM:
//		manDev.DisconnectAllDevices()
//		os.Exit(0)
//	}
//}
