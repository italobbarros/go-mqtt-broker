package models

type TopicResponse struct {
	Id        uint64      `gorm:"primaryKey;autoIncrement;uniqueIndex;column:Id"`
	Name      string      `gorm:"not null;unique;column:Name"`
	Config    TopicConfig `gorm:"embedded"`
	Container Container   `gorm:"embedded"`
	Times
}
