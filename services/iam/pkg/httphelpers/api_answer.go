package httphelpers

type AnswerError struct {
	Code string `json:"code"`
}

type ApiAnswer struct {
	Status  string       `json:"status"`
	Message string       `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
	Error   *AnswerError `json:"error,omitempty"`
}
