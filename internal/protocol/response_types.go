package protocol

type ResponseConnect struct {
	Id        string
	Timeout   int
	KeepAlive int16
	Clean     bool
	Username  string
	Password  string
}

type ResponsePublish struct {
	Identifier int
	Topic      string
	Body       []byte
	dutFlag    bool
	Retained   bool
	Qos        int
}
