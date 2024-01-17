package models

import "time"

type PublishResponse struct {
	Id              uint64    `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Payload         string    `gorm:"column:Payload"`
	Qos             int       `gorm:"column:Qos"`
	TopicName       string    `gorm:"not null;column:TopicName"`
	TopicRetained   bool      `gorm:"not null;column:TopicRetained"` // Indica se a mensagem é retida ou não
	Container       Container `gorm:"embedded;column:Container"`
	Timestamp       time.Time `gorm:"column:Timestamp" example:"2024-01-16T12:00:00Z"`
	NumberTimestamp int64     `gorm:"column:NumberTimestamp"`
}
type TopicResponse struct {
	Id      uint64          `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name    string          `gorm:"not null;unique;column:Name"`
	Publish PublishResponse `gorm:"embedded"`
	Times
}
