package models

import (
	"time"
)

type Container struct {
	Id   int    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name string `gorm:"not null;column:Name"`
}

type Times struct {
	Created time.Time  `gorm:"column:Created" example:"2024-01-16T12:00:00Z"`
	Updated time.Time  `gorm:"column:Updated" example:"2024-01-16T12:00:00Z"`
	Deleted *time.Time `gorm:"column:Deleted" example:"2024-01-16T12:45:00Z"`
}

type Topic struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	IdContainer uint64      `gorm:"column:IdContainer"`
	Container   Container   `gorm:"foreignKey:IdContainer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string      `gorm:"not null;unique;column:Name"`
	Config      TopicConfig `gorm:"embedded"`
	Times
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
