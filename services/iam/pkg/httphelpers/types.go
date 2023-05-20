package httphelpers

// ApiAnswer is the standard answer format for the API
type ApiAnswer struct {
	Status  string       `json:"status"`
	Message string       `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
	Error   *AnswerError `json:"error,omitempty"`
}

type AnswerError struct {
	Code string `json:"code"`
}
