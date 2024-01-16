package models

type ContainerRequest struct {
	Name string `gorm:"not null"`
}

type TopicConfig struct {
	Payload      string `gorm:"not null;column:Payload"`
	Qos          int    `gorm:"not null;column:Qos"`
	Retained     bool   `gorm:"not null;column:Retained"` // Indica se a mensagem é retida ou não
	SecurityRule string `gorm:"column:SecurityRule"`      // Regra de segurança aplicada ao tópico
}

type TopicRequest struct {
	IdContainer uint64
	Name        string      `gorm:"not null;column:Name"`
	Config      TopicConfig `gorm:"embedded"`
}

type PublicationRequest struct {
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	Payload     string    `gorm:"column:Payload"`
	Timestamp   int64     `gorm:"column:Timestamp"`
}

type SubscriptionRequest struct {
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	SessionId   string
}

type SessionRequest struct {
	IdContainer uint64 `gorm:"column:IdContainer"`
	ClientId    string `gorm:"column:ClientId"`
	KeepAlive   int16  `gorm:"column:KeepAlive"`
	Clean       bool   `gorm:"column:Clean"`
	Username    string `gorm:"column:Username"`
	Password    string `gorm:"column:Password"`
}
