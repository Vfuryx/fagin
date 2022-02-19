package enums

import (
	"fagin/pkg/enum"
)

// TimeFormat 枚举类型
type TimeFormat string

var _ enum.Enum = new(TimeFormat)

const (
	TimeFormatDef TimeFormat = "2006-01-02 15:04:05"
)

func (code TimeFormat) String() string {
	if code == TimeFormatDef {
		return "默认"
	}

	return ""
}

func (code TimeFormat) Get() string {
	return string(code)
}

func TimeFormatAll() map[string]string {
	return map[string]string{
		TimeFormatDef.Get(): TimeFormatDef.String(),
	}
}
