package api

type TreeNode struct {
	Name     string      `json:"name"`
	Topic    string      `json:"topic"`
	Children []*TreeNode `json:"children,omitempty"`
}

type TopicInfo struct {
	TopicName    string `json:"topicName"`
	Description  string `json:"description"`
	MessageCount int    `json:"messageCount"`
	Subscribers  int    `json:"subscribers"`
}

type SessionInfo struct {
	NumberSessions int `json:"numberSessions"`
}
