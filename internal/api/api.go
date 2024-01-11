package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	brokerMqtt "github.com/italobbarros/go-mqtt-broker/internal/broker"
)

// API representa a interface da API do servidor
type API struct {
	Broker *brokerMqtt.Broker
}

// Init inicializa a API com as rotas e inicia o servidor
func NewAPI(broker *brokerMqtt.Broker) *API {
	return &API{
		Broker: broker,
	}
}

func (a *API) Init() {
	r := mux.NewRouter()

	r.HandleFunc("/mqtt-tree", corsHandler(a.MqttTreeHandler)).Methods("GET")
	r.HandleFunc("/topic-info", corsHandler(a.TopicInfoHandler)).Methods("GET")

	http.Handle("/", r)

	// Inicie o servidor
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func (a *API) TopicInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Suponha que você esteja buscando informações do tópico de um banco de dados ou outro recurso.
	// Por agora, vamos simular algumas informações fictícias.
	topic := r.URL.Query().Get("topic")
	node := a.Broker.GetTopicNode(topic)
	if node == nil {
		http.Error(w, "topic not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(node)
}

func (a *API) MqttTreeHandler(w http.ResponseWriter, r *http.Request) {
	// Aqui, você pode usar a instância de Broker para obter a árvore MQTT.
	root := a.Broker.Root // Esta é uma suposição; você pode precisar implementar o método correspondente no seu Broker.

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

// Você também pode adicionar métodos adicionais à estrutura API conforme necessário.
