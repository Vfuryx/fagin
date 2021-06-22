package config

import (
	"embed"
	"fagin/resources"
	"fagin/resources/static"
	t "html/template"
	"time"
)

// 模版配置
type template struct {
	Public       string // 设置公开静态资源
	PublicRouter string // 设置公开静态资源路由地址
	Static       string // 设置静态资源
	StaticRouter string // 设置静态资源路由地址

	DelimitersL     string    // 模版定界符左
	DelimitersR     string    // 模版定界符右
	FuncMap         t.FuncMap // 模版函数
	StaticEmbed     embed.FS  // 静态资源（只读）
	TemplatesEmbed  embed.FS  // 模版静态资源（只读）
	TemplatePattern string    // 模版路径
}

var Template template

func init() {
	Template.Public = App.PublicPath
	Template.PublicRouter = "/public"
	Template.Static = App.StaticPath
	Template.StaticRouter = "/static"
	Template.StaticEmbed = static.Static

	Template.DelimitersL = "[["
	Template.DelimitersR = "]]"
	Template.FuncMap = FuncMap()
	Template.TemplatesEmbed = resources.Templates
	Template.TemplatePattern = "views/**/*"
}

func FuncMap() t.FuncMap {
	return t.FuncMap{
		"WebAsset": WebAsset,
		"HtmlSafe": HtmlSafe,
	}
}

func WebAsset(path string) string {
	if App.Env == "local" {
		return path
	}
	return CDN.URL + path + "?t=" + time.Now().Format("200612154")
}

func HtmlSafe(html string) t.HTML {
	return t.HTML(html)
}
