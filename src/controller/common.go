package controller

type Envelope[T any] struct {
	IsErro   bool   `json:"error,omitempty"`
	ErroCode string `json:"error-code,omitempty"`
	Message  string `json:"message"`
	Data     *[]T   `json:"data,omitempty"`
}
