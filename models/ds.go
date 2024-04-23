package models

type RequestID string

type Message struct {
	Text      string `json:"text"`
	RequestID string `json:"request_id"`
}

type HTTPResponse struct {
	Status      string `json:"status"`
	TimeElapsed int64  `json:"timeElapsed"`
	Message     any    `json:"response"`
	Error       string `json:"error,omitempty"`
}
