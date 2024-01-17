package models

import (
	"time"
)

type Times struct {
	Created time.Time `gorm:"column:Created" example:"2024-01-16T12:00:00Z"`
	Updated time.Time `gorm:"column:Updated" example:"2024-01-16T12:00:00Z"`
}
type Container struct {
	Id   int    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name string `gorm:"not null;column:Name"`
}

type Session struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64    `gorm:"column:IdContainer"`
	Container   Container `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientId    string    `gorm:"column:ClientId;unique"`
	KeepAlive   int16     `gorm:"column:KeepAlive"`
	Clean       bool      `gorm:"column:Clean"`
	Username    string    `gorm:"column:Username"`
	Password    string    `gorm:"column:Password"`
	Times
}

type Publish struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdSession uint64    `gorm:"column:IdSession"`
	Session   Session   `gorm:"foreignKey:IdSession;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TopicName string    `gorm:"not null;column:TopicName"`
	Payload   string    `gorm:"column:Payload"`
	Qos       int       `gorm:"column:Qos"`
	Timestamp time.Time `gorm:"column:Timestamp" example:"2024-01-16T12:00:00Z"`
}

type Topic struct {
	Id        uint64  `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdPublish uint64  `gorm:"not null;column:IdPublish"`
	Publish   Publish `gorm:"foreignKey:IdPublish"`
	Name      string  `gorm:"not null;unique;column:Name"`
	Retained  bool    `gorm:"not null;column:Retained"` // Indica se a mensagem é retida ou não
	Times
}

type Subscription struct {
	Id        uint64  `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdSession uint64  `gorm:"column:IdSession"`
	Session   Session `gorm:"foreignKey:IdSession;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdTopic   uint64  `gorm:"-;column:IdTopic"`
	Topic     Topic   `gorm:"foreignKey:IdTopic;"`
	Times
}
