package errno

type Errno interface {
	Get() int
	Error() string
}

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
