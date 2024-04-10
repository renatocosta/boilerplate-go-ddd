package http

type SelectLogFileRequest struct {
	Name string `json:"name"  binding:"required"`
}
