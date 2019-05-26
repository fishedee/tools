package exception

import (
	"fmt"
	"runtime"
	"strings"
)

// Exception Exception
type Exception struct {
	code    int
	message string
	stack   []string
	cause   interface{}
	isCrash bool
}

// NewException NewException
func NewException(code int, message string, args ...interface{}) *Exception {
	return newException(2, nil, false, code, message, args...)
}

func newException(stackBegin int, cause interface{}, isCrash bool, code int, message string, args ...interface{}) *Exception {
	if len(args) != 0 {
		message = fmt.Sprintf(message, args...)
	}

	stack := []string{}
	for i := stackBegin; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, fmt.Sprintf("%s:%d", file, line))
	}

	return &Exception{
		code:    code,
		message: message,
		stack:   stack,
		cause:   cause,
	}
}

// GetCode 获取Code
func (e *Exception) GetCode() int {
	return e.code
}

// GetMessage GetMessage
func (e *Exception) GetMessage() string {
	return e.message
}

// GetCause GetCause
func (e *Exception) GetCause() interface{} {
	return e.cause
}

// IsCrash IsCrash
func (e *Exception) IsCrash() bool {
	return e.isCrash
}

// GetStackTrace GetStackTrace
func (e *Exception) GetStackTrace() string {
	return strings.Join(e.stack, "\n")
}

// GetStackTraceLine GetStackTraceLine
func (e *Exception) GetStackTraceLine(i int) string {
	return e.stack[i]
}

// Error Error
func (e Exception) Error() string {
	return fmt.Sprintf("[Code:%d] [Message:%s] [Stack:%s]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
}

// Throw Throw
func Throw(code int, message string, args ...interface{}) {
	exception := newException(2, nil, false, code, message, args...)

	panic(exception)
}

// CatchCrash CatchCrash
func CatchCrash(handler func(Exception)) {
	err := recover()
	if err != nil {
		exception, isException := err.(*Exception)
		if isException {
			handler(*exception)
		} else {
			exception := newException(4, err, true, 1, fmt.Sprint(err))
			handler(*exception)
		}
	}
}

// Catch Catch
func Catch(handler func(Exception)) {
	err := recover()
	if err != nil {
		exception, isException := err.(*Exception)
		if isException {
			if exception.IsCrash() == false {
				handler(*exception)
			} else {
				panic(exception)
			}
		} else {
			exception := newException(4, err, true, 1, fmt.Sprint(err))
			panic(exception)
		}
	}
}
