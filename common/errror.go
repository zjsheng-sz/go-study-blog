package common

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppErr(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrNotFound     = NewAppErr(404, "souce not found", nil)
	ErrUnauthorized = NewAppErr(401, "not auth", nil)
	ErrInternal     = NewAppErr(500, "service internal error", nil)
)
