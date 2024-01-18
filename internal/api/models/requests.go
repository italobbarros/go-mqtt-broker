package models

type ContainerRequest struct {
	Name string `gorm:"not null"`
}

type TopicRequest struct {
	IdPublish uint64 `gorm:"not null;column:IdPublish"`
}

type PublishRequest struct {
	ClientIdSession string `gorm:"column:ClientIdSession;not null;"`
	Payload         string `gorm:"column:Payload"`
	Qos             int    `gorm:"column:Qos"`
	TopicName       string `gorm:"not null;column:TopicName"`
	TopicRetained   bool   `gorm:"not null;column:TopicRetained"` // Indica se a mensagem é retida ou não
}

type SubscriptionRequest struct {
	IdSession uint64 `gorm:"column:IdSession"`
	IdTopic   uint64 `gorm:"column:IdTopic"`
}

type SessionRequest struct {
	IdContainer uint64 `gorm:"column:IdContainer"`
	ClientId    string `gorm:"column:ClientId"`
	KeepAlive   int16  `gorm:"column:KeepAlive"`
	Clean       bool   `gorm:"column:Clean"`
	Username    string `gorm:"column:Username"`
	Password    string `gorm:"column:Password"`
}

type SessionUpdateRequest struct {
	IdContainer uint64 `gorm:"column:IdContainer"`
	KeepAlive   int16  `gorm:"column:KeepAlive"`
	Clean       bool   `gorm:"column:Clean"`
	Username    string `gorm:"column:Username"`
	Password    string `gorm:"column:Password"`
}
