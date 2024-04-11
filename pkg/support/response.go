package support

type Response[T any] struct {
	StatusCode int      `json:"-"`
	Data       *T       `json:"data,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}
