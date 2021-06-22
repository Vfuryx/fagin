package time_format

import "fagin/pkg/constant"

const (
	Def string = "2006-01-02 15:04:05"
)

type timeFormat struct {
	Def string
	constant.Constant
}

func TimeFormat() *timeFormat {
	mt := timeFormat{
		Def: Def,
	}
	mt.Constant.C = &mt
	return &mt
}

func (s *timeFormat) All() map[interface{}]string {
	return map[interface{}]string{
		s.Def: "默认",
	}
}
