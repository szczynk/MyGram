package delivery

type ErrorResponse struct {
	Error   string `json:"error" example:"error"`
	Message string `json:"message" example:"message"`
}
