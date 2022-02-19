package errno

// Errno 错误码
type Errno interface {
	Get() int
	Error() string
}

// Decode 解码
func Decode(err error) (int, string) {
	if err == nil {
		return 1, "未知错误"
	}

	e, ok := err.(Errno)
	if !ok {
		return 1, err.Error()
	}

	return e.Get(), e.Error()
}
