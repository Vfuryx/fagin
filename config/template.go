package config

import (
	"embed"
	"fadmin/resources"
	"fadmin/resources/static"
	t "html/template"
	"sync/atomic"
	"time"
)

var templateConfig atomic.Value

func Template() *TemplateConfig {
	if c, ok := templateConfig.Load().(*TemplateConfig); ok {
		return c
	}

	return &TemplateConfig{}
}

func templateConfigInit() {
	c := &TemplateConfig{
		public:       App().PublicPath(),
		publicRouter: "/public",
		static:       App().StaticPath(),
		staticRouter: "/static",

		delimitersL:     "[[",
		delimitersR:     "]]",
		funcMap:         FuncMap(),
		staticEmbed:     static.Static,
		templatesEmbed:  resources.Templates,
		templatePattern: "views/**/*",
	}

	templateConfig.Store(c)
}

func FuncMap() t.FuncMap {
	return t.FuncMap{
		"WebAsset": WebAsset,
		"HtmlSafe": HTMLSafe,
	}
}

func WebAsset(path string) string {
	if App().Env() == "local" {
		return path
	}

	return CDN().URL() + path + "?t=" + time.Now().Format("200612154")
}

func HTMLSafe(html interface{}) t.HTML {
	if v, ok := html.(t.HTML); ok {
		return v
	}

	return ""
}

// TemplateConfig 模版配置
type TemplateConfig struct {
	public       string // 设置公开静态资源
	publicRouter string // 设置公开静态资源路由地址
	static       string // 设置静态资源
	staticRouter string // 设置静态资源路由地址

	delimitersL     string    // 模版定界符左
	delimitersR     string    // 模版定界符右
	funcMap         t.FuncMap // 模版函数
	staticEmbed     embed.FS  // 静态资源（只读）
	templatesEmbed  embed.FS  // 模版静态资源（只读）
	templatePattern string    // 模版路径
}

func (tpl *TemplateConfig) Public() string {
	return tpl.public
}

func (tpl *TemplateConfig) PublicRouter() string {
	return tpl.publicRouter
}

func (tpl *TemplateConfig) Static() string {
	return tpl.static
}

func (tpl *TemplateConfig) StaticRouter() string {
	return tpl.staticRouter
}

func (tpl *TemplateConfig) DelimitersL() string {
	return tpl.delimitersL
}

func (tpl *TemplateConfig) DelimitersR() string {
	return tpl.delimitersR
}

func (tpl *TemplateConfig) FuncMap() t.FuncMap {
	return tpl.funcMap
}

func (tpl *TemplateConfig) StaticEmbed() embed.FS {
	return tpl.staticEmbed
}

func (tpl *TemplateConfig) TemplatesEmbed() embed.FS {
	return tpl.templatesEmbed
}

func (tpl *TemplateConfig) TemplatePattern() string {
	return tpl.templatePattern
}
