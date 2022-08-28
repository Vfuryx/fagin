package enums

import (
	"fagin/pkg/enum"
)

// Status 枚举类型
type Status uint8

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

func (code Status) Get() uint8 {
	return uint8(code)
}

func AllStatus() map[uint8]string {
	return map[uint8]string{
		StatusActive.Get():  StatusActive.String(),
		StatusDisable.Get(): StatusDisable.String(),
	}
}
