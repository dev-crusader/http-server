package models

type RequestID string

type Message struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	RequestID string `json:"request_id"`
}

type HTTPResponse struct {
	Status      string  `json:"status"`
	TimeElapsed int64   `json:"timeElapsed"`
	Message     Message `json:"response"`
	Error       string  `json:"error,omitempty"`
}

type MessageBody struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}
