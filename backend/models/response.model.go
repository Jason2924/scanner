package models

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(message string, data any) *Response {
	return &Response{
		Message: message,
		Data:    data,
	}
}
