package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/mqtt-tree", corsHandler(mqttTreeHandler)).Methods("GET")
	r.HandleFunc("/topic-info", corsHandler(topicInfoHandler)).Methods("GET")

	http.Handle("/", r)

	// Inicie o servidor
	http.ListenAndServe(":8080", nil)
}

func topicInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Suponha que você esteja buscando informações do tópico de um banco de dados ou outro recurso.
	// Por agora, vamos simular algumas informações fictícias.
	topic := r.URL.Query().Get("topic")
	topicInfo := TopicInfo{
		TopicName:    topic,
		Description:  "Esta é uma descrição para o tópico de exemplo.",
		MessageCount: 1000,
		Subscribers:  50,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topicInfo)
}

func mqttTreeHandler(w http.ResponseWriter, r *http.Request) {
	// Simulação de uma árvore de tópicos MQTT simples
	root := &TreeNode{
		Name:  "container1",
		Topic: "container1/*",
		Children: []*TreeNode{
			{
				Name:  "client1",
				Topic: "container1/client1/*",
				Children: []*TreeNode{
					{
						Name:  "teste",
						Topic: "container1/client1/teste/*",

						Children: []*TreeNode{
							{
								Name:  "io1",
								Topic: "container1/client1/teste/io1",
							},
							{
								Name:  "io1",
								Topic: "container1/client1/teste/io2",
							},
						},
					},
				},
			},
			{
				Name:  "client2",
				Topic: "container1/client1/*",
				Children: []*TreeNode{
					{
						Name:  "io1",
						Topic: "container1/client2/teste/io1",
					},
					{
						Name:  "io1",
						Topic: "container1/client2/teste/io2",
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}
