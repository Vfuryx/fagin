package status

import "fagin/pkg/constant"

// Disable 关闭
const Disable = 0

// Active 开启
const Active = 1

type status struct {
	Disable int
	Active  int
	constant.Constant
}

func Status() *status {
	s := status{
		Disable: Disable,
		Active:  Active,
	}
	s.Constant.C = &s
	return &s
}

func (s *status) All() map[interface{}]string {
	return map[interface{}]string{
		s.Active:  "激活",
		s.Disable: "关闭",
	}
}
