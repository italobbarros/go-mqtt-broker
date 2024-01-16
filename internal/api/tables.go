package api

import (
	"time"
)

type ContainerPost struct {
	Name string `gorm:"not null"`
}

type Container struct {
	Id   int    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name string `gorm:"not null;column:Name"`
}

type TopicConfig struct {
	Payload      string `gorm:"not null;column:Payload"`
	Qos          int    `gorm:"not null;column:Qos"`
	Retained     bool   `gorm:"not null;column:Retained"` // Indica se a mensagem é retida ou não
	SecurityRule string `gorm:"column:SecurityRule"`      // Regra de segurança aplicada ao tópico
}

type Times struct {
	Created time.Time  `gorm:"column:Created" example:"2024-01-16T12:00:00Z"`
	Updated time.Time  `gorm:"column:Updated" example:"2024-01-16T12:00:00Z"`
	Deleted *time.Time `gorm:"column:Deleted" example:"2024-01-16T12:45:00Z"`
}

type TopicRequest struct {
	IdContainer uint64
	Name        string      `gorm:"not null;column:Name"`
	Config      TopicConfig `gorm:"embedded"`
}

type TopicResponse struct {
	Id        uint64      `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name      string      `gorm:"not null;unique;column:Name"`
	Config    TopicConfig `gorm:"embedded"`
	Container Container   `gorm:"embedded"`
	Times
}

type Topic struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64      `gorm:"column:IdContainer"`
	Container   Container   `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string      `gorm:"not null;unique;column:Name"`
	Config      TopicConfig `gorm:"embedded"`
	Times
}

type PublicationRequest struct {
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	Payload     string    `gorm:"column:Payload"`
	Timestamp   int64     `gorm:"column:Timestamp"`
}

type Publication struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	Payload     string    `gorm:"column:Payload"`
	Timestamp   int64     `gorm:"column:Timestamp"`
	Times
}

type SubscriptionRequest struct {
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	SessionId   string
}

type Subscription struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic     uint64    `gorm:"-;column:IdTopic"`
	Topic       Topic     `gorm:"foreignKey:IdTopic;"`
	IdSession   uint64    `gorm:"-;column:IdSession"`
	Session     Session   `gorm:"foreignKey:IdSession"`
	Times
}

type SessionRequest struct {
	IdContainer uint64 `gorm:"column:IdContainer"`
	ClientId    string `gorm:"column:ClientId"`
	KeepAlive   int16  `gorm:"column:KeepAlive"`
	Clean       bool   `gorm:"column:Clean"`
	Username    string `gorm:"column:Username"`
	Password    string `gorm:"column:Password"`
}

type Session struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientId    string
	KeepAlive   int16
	Clean       bool
	Username    string
	Password    string
	Times
}
