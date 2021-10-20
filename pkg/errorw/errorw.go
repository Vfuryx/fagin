package errorw

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// ErrorStack 堆栈
var ErrorStack = errors.New("error stack")

// New 实例化
func New(text string) error {
	return WithStack(errors.New(text))
}

// UP WithStack 别名
func UP(err error) error {
	return WithStack(err)
}

// WithStack 带上堆栈信息
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, ErrorStack) {
		return err
	}
	return fmt.Errorf(err.Error()+": \n%w", fmt.Errorf(ErrorStack.Error()+": \n%w", errors.New(string(debug.Stack()))))
}

// Wrap Wrap
func Wrap(err, wrap error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf(err.Error()+": \n%w", WithStack(wrap))
}
