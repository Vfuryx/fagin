package constant

type constant interface {
	All() map[interface{}]string
	Get(code interface{}) string
	IsExist(code interface{}) bool
}

type Constant struct {
	C constant
}

func (c *Constant) Get(code interface{}) string {
	if c.C.IsExist(code) {
		return c.C.All()[code]
	}
	return ""
}

func (c *Constant) IsExist(code interface{}) bool {
	_, ok := c.C.All()[code]
	return ok
}
