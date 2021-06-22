package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type adminCasbin struct {
	cache.SCache
}

func NewAdminCasbin(f cache.GetterFunc) *adminCasbin {
	var c = new(adminCasbin)
	c.Prefix = "admin::casbin::"
	c.LifeTime = time.Minute
	c.Content = f
	c.SetFunc(c)
	return c
}

// Key 获取键名称
func (c *adminCasbin) Key(value string) string {
	return c.Prefix + value
}

// Lift 默认存在时间
func (c *adminCasbin) Lift() time.Duration {
	return c.LifeTime
}
