package error_handler

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHTTPError(code int, msg string) *HTTPError {
	he := &HTTPError{Code: code, Message: msg}
	return he
}

func (he HTTPError) Error() string {
	return he.Message
}
