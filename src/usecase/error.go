package usecase

type ApplicationError struct {
	Code       int    `json:"code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	OriginErr  string `json:"error"`
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) AddError(err error) error {
	return NewApplicationError(e.Code, e.StatusCode, e.Message, err)
}

func NewApplicationError(code int, statusCode int, message string, err error) *ApplicationError {
	return &ApplicationError{
		Code:       code,
		StatusCode: statusCode,
		Message:    message,
		OriginErr:  err.Error(),
	}
}
