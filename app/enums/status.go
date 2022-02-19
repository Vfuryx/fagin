package enums

import (
	"fagin/pkg/enum"
)

// Status 枚举类型
type Status int

var _ enum.Enum = new(Status)

// Disable
const (
	StatusDisable Status = 0 // 关闭
	StatusActive  Status = 1 // 开启
)

func (code Status) String() string {
	switch code {
	case StatusDisable:
		return "关闭"
	case StatusActive:
		return "开启"
	}

	return ""
}

func (code Status) Get() int {
	return int(code)
}

func AllStatus() map[int]string {
	return map[int]string{
		StatusActive.Get():  StatusActive.String(),
		StatusDisable.Get(): StatusDisable.String(),
	}
}
