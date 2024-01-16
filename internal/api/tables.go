package api

import "gorm.io/gorm"

type ContainerPost struct {
	ContainerName string `gorm:"not null"`
}

type Container struct {
	ID            int    `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ContainerName string `gorm:"not null"`
}

type TopicConfig struct {
	Payload      string
	Qos          int
	Retained     bool   // Indica se a mensagem é retida ou não
	SecurityRule string // Regra de segurança aplicada ao tópico
}

type TopicPost struct {
	ContainerID int
	Container   Container   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicName   string      `gorm:"not null"`
	TopicConfig TopicConfig `gorm:"embedded"`
	gorm.Model
}

type Topic struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ContainerID int
	Container   Container   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicName   string      `gorm:"not null"`
	TopicConfig TopicConfig `gorm:"embedded"`
	gorm.Model
}

type SubscriptionPost struct {
	ContainerID int
	Container   Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicID     uint64
	Topic       Topic
	SessionID   string
	gorm.Model
}

type Subscription struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ContainerID int
	Container   Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicID     uint64
	Topic       Topic
	SessionID   string
	gorm.Model
}

type PublicationPost struct {
	ContainerID    uint64
	Container      Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicID        uint64
	Topic          Topic
	MessagePayload string
	Timestamp      int64
	gorm.Model
}

type Publication struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ContainerID    uint64
	Container      Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicID        uint64
	Topic          Topic
	MessagePayload string
	Timestamp      int64
	gorm.Model
}

type SessionPost struct {
	ContainerID uint64
	Container   Container `gorm:"foreignKey:ContainerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientID    string
	KeepAlive   int16
	Clean       bool
	Username    string
	Password    string
	gorm.Model
}

type Session struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ContainerID uint64
	Container   Container `gorm:"foreignKey:ContainerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientID    string
	KeepAlive   int16
	Clean       bool
	Username    string
	Password    string
	gorm.Model
}
