package resources

import "embed"

// Templates 模版资源（只读）
//go:embed views/*
var Templates embed.FS
