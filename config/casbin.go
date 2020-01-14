package config

import (
	"fagin/pkg/conf"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"strconv"
)

// Casbin 模型配置
var Casbin model.Model

func init() {
	Casbin = casbin.NewModel()

	m := conf.GetStringMapStringSlice("casbin.model", map[string][]string{})

	mLen := len(m)
	for i := 0; i < mLen; i++ {
		c := m[strconv.Itoa(i)]
		Casbin.AddDef(c[0], c[1], c[2])
	}

	// rbac_model.conf 参考 https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf
	//Casbin.AddDef("r", "r", "sub, obj, act")
	//Casbin.AddDef("p", "p", "sub, obj, act")
	//Casbin.AddDef("g", "g", "_, _")
	//Casbin.AddDef("e", "e", "some(where (p.eft == allow))")
	//Casbin.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")

	//// RESTful 模型 参考 https://casbin.org/docs/zh-CN/supported-models
	//Casbin.AddDef("r", "r", "sub, obj, act")
	//Casbin.AddDef("p", "p", "sub, obj, act")
	//Casbin.AddDef("e", "e", "some(where (p.eft == allow))")
	//Casbin.AddDef("m", "m", "r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")
}
