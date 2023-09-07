package exception

const (
	ErrKeyInvalidRequest      = "error.invalidRequest"
	ErrKeyInternalServerError = "error.internalServerError"
)

type Error struct {
	message string
	code    string
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Error() string {
	return e.message
}

func NewError(code string, message string) *Error {
	return &Error{
		message: message,
		code:    code,
	}
}

func NewInvalidRequestError(message string) *Error {
	return &Error{
		message: message,
		code:    ErrKeyInvalidRequest,
	}
}
