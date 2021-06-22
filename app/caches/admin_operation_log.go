package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 后台操作缓存管理
type adminOperationLog struct {
	cache.SCache
}

func NewAdminOperationLog(f cache.GetterFunc) *adminOperationLog {
	var c = new(adminOperationLog)
	c.Prefix = "admin::menu::"
	c.LifeTime = 30 * time.Second
	c.Content = f
	c.SetFunc(c)
	return c
}

// Key 获取键名称
func (c *adminOperationLog) Key(value string) string {
	return c.Prefix + value
}

// Lift 默认存在时间
func (c *adminOperationLog) Lift() time.Duration {
	return c.LifeTime
}
