package enum

type Enum interface {
	String() string
}

func IsExist(code Enum) bool {
	return code.String() != ""
}
