package boom

import "net/http"

func BadImplementation() *ErrorMessage {
	return &ErrorMessage{
		Code:    http.StatusInternalServerError,
		Error:   "Internal Server Error",
		Message: "An internal server error occurred",
	}
}
