package types

const (
	SUCCESS = "success"
	ERROR   = "error"
)

type HttpResponse[T any] struct {
	Status     string
	StatusCode int
	Message    string
	Data       T
}
