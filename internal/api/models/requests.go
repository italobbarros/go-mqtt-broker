package models

type ContainerRequest struct {
	Name string `gorm:"not null"`
}

type TopicRequest struct {
	IdPublish uint64 `gorm:"not null;column:IdPublish"`
}

type PublishRequest struct {
	IdContainer   uint64 `gorm:"column:IdContainer"`
	Payload       string `gorm:"column:Payload"`
	Qos           int    `gorm:"column:Qos"`
	TopicName     string `gorm:"not null;column:TopicName"`
	TopicRetained bool   `gorm:"not null;column:TopicRetained"` // Indica se a mensagem é retida ou não
	//Timestamp       time.Time `gorm:"column:Timestamp" example:"2024-01-16T12:00:00Z"`
	//NumberTimestamp int64     `gorm:"column:NumberTimestamp"`
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
