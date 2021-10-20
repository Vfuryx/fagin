package config

import (
	"embed"
	"fagin/resources"
	"fagin/resources/static"
	t "html/template"
	"time"
)

// TemplateConfig 模版配置
type TemplateConfig struct {
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

var templateConfig = new(TemplateConfig)

func Template() TemplateConfig {
	return *templateConfig
}

func (template *TemplateConfig) init() {
	template.Public = App().PublicPath
	template.PublicRouter = "/public"
	template.Static = App().StaticPath
	template.StaticRouter = "/static"
	template.StaticEmbed = static.Static

	template.DelimitersL = "[["
	template.DelimitersR = "]]"
	template.FuncMap = FuncMap()
	template.TemplatesEmbed = resources.Templates
	template.TemplatePattern = "views/**/*"
}

func FuncMap() t.FuncMap {
	return t.FuncMap{
		"WebAsset": WebAsset,
		"HtmlSafe": HtmlSafe,
	}
}

func WebAsset(path string) string {
	if App().Env == "local" {
		return path
	}
	return cdnConfig.URL + path + "?t=" + time.Now().Format("200612154")
}

func HtmlSafe(html string) t.HTML {
	return t.HTML(html)
}
