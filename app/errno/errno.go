package errno

import (
	"fmt"
)

type Errno struct {
	Code    int
	Message string
}

func (errno *Errno) Error() string {
	return errno.Message
}

var (
	OK = &Errno{
		Code:    0,
		Message: "OK",
	}
	InternalServerError = &Errno{
		Code:    500,
		Message: "Internal server exception",
	}
)

// Err represents an exception
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, exception: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
