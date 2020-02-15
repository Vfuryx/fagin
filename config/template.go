package config

import (
	t "html/template"
	"time"
)

// 模版配置
type template struct {
	FuncMap      t.FuncMap // 模版函数
	LoadHTMLGlob string    // 设置模版目录
	Static       string    // 设置静态资源
	StaticRouter string    // 设置静态资源路由地址
	DelimitersL  string    // 模版定界符左
	DelimitersR  string    // 模版定界符右
}

var Template template

func init() {
	Template.FuncMap = FuncMap()
	Template.LoadHTMLGlob = App.ResourcePath + "/views/**/*"
	Template.Static = App.PublicPath
	Template.StaticRouter = "/public"
	Template.DelimitersL = "[["
	Template.DelimitersR = "]]"
}

func FuncMap() t.FuncMap {
	return t.FuncMap{
		"WebAsset": WebAsset,
	}
}

func WebAsset(path string) string {
	if App.Env == "local" {
		return path
	}
	return CDN.URL + path + "?t=" + time.Now().Format("200612154")
}
