package boom

import "net/http"

func BadRequest(message string, data interface{}) *ErrorMessage {
	return &ErrorMessage{
		Code:    http.StatusBadRequest,
		Error:   "Bad Request",
		Message: message,
		Data:    data,
	}
}

func Unauthorized(message string) *ErrorMessage {
	// TODO: Support other params like https://github.com/hapijs/boom#boomunauthorizedmessage-scheme-attributes
	return &ErrorMessage{
		Code:    http.StatusUnauthorized,
		Error:   "Bad Request",
		Message: message,
	}
}

func PaymentRequired(message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    http.StatusPaymentRequired,
		Error:   "Payment Required",
		Message: message,
	}
}

func Forbidden(message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    http.StatusForbidden,
		Error:   "Forbidden",
		Message: message,
	}
}

func NotFound(message string) *ErrorMessage {
	return &ErrorMessage{
		Code:    http.StatusNotFound,
		Error:   "Not Found",
		Message: message,
	}
}
