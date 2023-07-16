package errorw

import (
	"errors"
	"runtime/debug"
	"strings"
	"sync"
)

var strBufPool = sync.Pool{
	New: func() interface{} { return new(strings.Builder) },
}

type errorw struct {
	error
	stack []byte
}

func (err errorw) Error() string {
	var b = strBufPool.Get().(*strings.Builder)
	defer func() {
		b.Reset()
		strBufPool.Put(b)
	}()
	b.WriteString(err.error.Error())
	b.WriteString("\nerror stack \n")
	b.Write(err.stack)

	return b.String()
}

// New 实例化
func New(text string) error {
	return errorw{
		error: errors.New(text),
		stack: debug.Stack(),
	}
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

	switch err.(type) {
	case errorw:
		return err
	default:
		return New(err.Error())
	}
}

// Wrap 包装
func Wrap(err, wrap error) error {
	var b = strBufPool.Get().(*strings.Builder)
	defer func() {
		b.Reset()
		strBufPool.Put(b)
	}()

	if err == nil {
		return nil
	}

	if e, ok := err.(errorw); ok {
		b.WriteString(e.error.Error())
	} else {
		b.WriteString(err.Error())
	}

	b.WriteByte('\n')

	if e, ok := wrap.(errorw); ok {
		b.WriteString(e.error.Error())

		return errorw{
			error: errors.New(b.String()),
			stack: e.stack,
		}
	} else {
		b.WriteString(wrap.Error())
		return New(b.String())
	}
}

func Stack(err error) []byte {
	if r, ok := err.(errorw); ok {
		return r.stack
	}

	return []byte{}
}
