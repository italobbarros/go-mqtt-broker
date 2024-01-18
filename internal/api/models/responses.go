package models

import "time"

type GenericResponse struct {
	Detail string `gorm:"column:detail"`
}

type PublishResponse struct {
	Id              uint64              `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	ClientIdSession string              `gorm:"column:ClientIdSession;not null;"`
	Payload         string              `gorm:"column:Payload"`
	Qos             int                 `gorm:"column:Qos"`
	TopicName       string              `gorm:"not null;column:TopicName"`
	TopicRetained   bool                `gorm:"not null;column:TopicRetained"` // Indica se a mensagem é retida ou não
	Session         SessionInfoResponse `gorm:"embedded;column:Session"`
	Timestamp       time.Time           `gorm:"column:Timestamp" example:"2024-01-16T12:00:00Z"`
	NumberTimestamp int64               `gorm:"column:NumberTimestamp"`
}
type TopicResponse struct {
	Id      uint64          `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name    string          `gorm:"not null;unique;column:Name"`
	Publish PublishResponse `gorm:"embedded"`
	Times
}

type SessionResponse struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Container Container `gorm:"embedded;column:Container"`
	ClientId  string
	KeepAlive int16
	Clean     bool
	Username  string
	Password  string
	Times
}

type SessionInfoResponse struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Container Container `gorm:"embedded;column:Container"`
	ClientId  string
	KeepAlive int16
	Clean     bool
}

type ContainersInfoResponse struct {
	Id               uint64 `gorm:"column:Id"`
	CountSession     uint64 `gorm:"column:CountSession"`
	CountPublishers  uint64 `gorm:"column:CountPublishers"`
	CountSubscribers uint64 `gorm:"column:CountSubscribers"`
}
