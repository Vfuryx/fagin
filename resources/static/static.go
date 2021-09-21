package static

import "embed"

// Static 静态资源（只读）
//go:embed admin web
var Static embed.FS
