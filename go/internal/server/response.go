package server

// ErrorResponse - generic error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
