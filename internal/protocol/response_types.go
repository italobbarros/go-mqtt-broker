package protocol

type ResponseConnect struct {
	Id        string
	Timeout   int
	keepAlive int16
	Clean     bool
	username  string
	password  string
}
