package errorx

type CodeError struct {
	Code    int
	Message string
}

func NewCodeError(code int, message string) *CodeError {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) GetCode() int {
	if e == nil {
		return 0
	}
	return e.Code
}

func (e *CodeError) GetMessage() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func ParseError(err error) *CodeError {
	codeError, ok := err.(*CodeError)
	if ok {
		return codeError
	}
	return &CodeError{99999, "unknown error"}
}
