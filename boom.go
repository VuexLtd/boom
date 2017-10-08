package boom

type ErrorMessage struct {
	Code    int         `json:"statusCode"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
