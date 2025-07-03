package apperrors

type ErrorType string

const (
	ValidationError   ErrorType = "VALIDATION"
	NotFoundError     ErrorType = "NOT_FOUND"
	TimeoutError      ErrorType = "TIMEOUT"
	UnauthorizedError ErrorType = "UNAUTHORIZED"
	ForbiddenError    ErrorType = "FORBIDDEN"
	ConflictError     ErrorType = "CONFLICT"
	ExternalAPIError  ErrorType = "EXTERNAL_API"
	InternalError     ErrorType = "INTERNAL"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return string(e.Type) + ": " + e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(errType ErrorType, msg string, err error) *AppError {
	return &AppError{
		Type:    errType,
		Message: msg,
		Err:     err,
	}
}
