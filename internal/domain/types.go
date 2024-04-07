package domain

type OptionFlag uint8

const (
	OptionFlagNotDefined OptionFlag = iota
	OptionFlagTrue
	OptionFlagFalse
)

// Error handling types
type ErrorType uint8

const (
	ErrNotFound ErrorType = iota
	ErrInvalid
)

type AppError struct {
	Message string
	Type    ErrorType
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Type:    ErrNotFound,
		Message: message,
	}
}

func NewInvalidError(message string) error {
	return AppError{
		Type:    ErrInvalid,
		Message: message,
	}
}
