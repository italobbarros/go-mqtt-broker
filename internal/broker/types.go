package broker

import (
	"sync"
	"time"

	"github.com/italobbarros/go-mqtt-broker/internal/protocol"
	connection "github.com/italobbarros/go-mqtt-broker/pkg/connection"
	"github.com/italobbarros/go-mqtt-broker/pkg/logger"
)

// Broker representa a entidade do corretor MQTT
type BrokerConfigs struct {
	Name           string
	Address        string
	TypeConnection connection.TypeConnection
}

type Broker struct {
	logger    *logger.Logger
	Root      *TopicNode
	SessionMg *SessionManager
	server    connection.ServerInterface
}

// Topic.go
type TopicNode struct {
	Name         string                       `json:"name"`
	Topic        string                       `json:"topic"`
	TopicConfig  *TopicConfig                 `json:"topicCfg,omitempty"`
	Subscribers  map[string]*SubscriberConfig `json:"subscribers"` // Lista de sub-tópicos associados a este tópico
	MessageCount int                          `json:"messageCount"`
	Children     []*TopicNode                 `json:"children,omitempty"`
}

type TopicConfig struct {
	Payload      string `json:"payload"`
	Qos          int    `json:"qos"`
	Retained     bool   `json:"retained"`     // Indica se a mensagem é retida ou não
	SecurityRule string `json:"securityRule"` // Regra de segurança aplicada ao tópico
}

type SubscriberConfig struct {
	Identifier []byte
	Qos        int
	session    *Session
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
	prot      *protocol.MqttProtocol
	config    *SessionConfig
	top       *Session // Ponteiro para o nó anterior na lista
	bottom    *Session // Ponteiro para o próximo nó na lista
	logger    *logger.Logger
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
