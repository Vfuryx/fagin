package static

import "embed"

// Static 静态资源（只读）
//go:embed admin web
var Static embed.FS

// AdminHTML 后台首页（只读）
//go:embed admin/index.html
var AdminHTML []byte
