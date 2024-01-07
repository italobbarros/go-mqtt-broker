package broker

import (
	"sync"
	"time"

	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
)

// Broker representa a entidade do corretor MQTT
type BrokerConfigs struct {
	Name           string
	Address        string
	TypeConnection connection.TypeConnection
}

type Broker struct {
	Root      *TopicNode
	SessionMg *SessionManager
	server    connection.ServerInterface
}

// Topic.go
type TopicNode struct {
	Name         string       `json:"name"`
	Topic        string       `json:"topic"`
	TopicConfig  *TopicConfig `json:"topicCfg,omitempty"`
	MessageCount int          `json:"messageCount"`
	Children     []*TopicNode `json:"children,omitempty"`
}

type TopicConfig struct {
	Payload      string                           `json:"payload"`
	Qos          int                              `json:"qos"`
	Retained     bool                             `json:"retained"`     // Indica se a mensagem é retida ou não
	Subscribers  []connection.ConnectionInterface `json:"subscribers"`  // Lista de sub-tópicos associados a este tópico
	SecurityRule string                           `json:"securityRule"` // Regra de segurança aplicada ao tópico
}

// session.go
type SessionConfig struct {
	Id        string
	KeepAlive int16
	Clean     bool
	username  string
	password  string
}

// Session representa uma sessão MQTT
type Session struct {
	Timestamp time.Time
	partition int
	config    *SessionConfig
	top       *Session // Ponteiro para o nó anterior na lista
	bottom    *Session // Ponteiro para o próximo nó na lista
}

type SessionPartition struct {
	head *Session // Ponteiro para o primeiro nó da lista
	tail *Session // Ponteiro para o ultimo nó da lista
}

// SessionManager gerencia sessões MQTT
type SessionManager struct {
	sessionMap    map[string]*Session // Mapa para acessar sessões por ID
	partitionMap  map[int16]*SessionPartition
	lockSession   sync.Mutex
	lockPartition sync.Mutex
}
