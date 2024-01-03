package broker

type TopicNode struct {
	Name        string       `json:"name"`
	Topic       string       `json:"topic"`
	TopicConfig *TopicConfig `json:"topicCfg,omitempty"`
	Children    []*TopicNode `json:"children,omitempty"`
}

type TopicConfig struct {
	QoS          int      // Nível de Qualidade de Serviço (0, 1 ou 2)
	Retained     bool     // Indica se a mensagem é retida ou não
	Subscribers  []string // Lista de sub-tópicos associados a este tópico
	SecurityRule string   // Regra de segurança aplicada ao tópico
}

type TopicInfo struct {
	TopicName    string `json:"topicName"`
	Description  string `json:"description"`
	MessageCount int    `json:"messageCount"`
	Subscribers  int    `json:"subscribers"`
}

// Broker representa a entidade do corretor MQTT
type Broker struct {
	Root *TopicNode
}
