package client

// MessageText for Pipeline Notification
type MessageText struct {
	Text string `json:"text"`
}

// Message for Pipeline Notification
type Message struct {
	Complete MessageText `json:"pipeline.complete"`
	Failed   MessageText `json:"pipeline.failed"`
	Starting MessageText `json:"pipeline.starting"`
}
