package cache

import (
	"fagin/app/errno"
	"time"
)

// 缓存接口
type oCache interface {
	Lift() time.Duration
	Key(value string) string
	GetContent(value string) (string, error)
}

type iCache interface {
	Get(key string) (string, error)
	Exists(key string) (int64, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	Remove(key string) (int64, error)
}

func cache() (iCache, error) {
	if Redis == nil {
		return nil, errno.Api.ErrCacheIsClose
	}
	return Redis, nil
}

type SCache struct {
	Prefix   string
	LifeTime time.Duration
	oCache
}

var _ iCache = &SCache{}

// 载入多态
func (sc *SCache) SetFunc(o oCache) {
	sc.oCache = o
}

// 是键是否存在
func (sc *SCache) Exists(value string) (int64, error) {
	c, err := cache()
	if err != nil {
		return 0, err
	}
	key := sc.Key(value)
	return c.Exists(key)
}

// 根据键获取数据
func (sc *SCache) Get(value string) (string, error) {
	c, err := cache()
	if err != nil {
		return "", err
	}
	key := sc.Key(value)
	ok, err := c.Exists(key)
	if err != nil {
		return "", err
	}
	// 存在
	if ok > 0 {
		return c.Get(key)
	}
	// 不存在
	content, err := sc.GetContent(value)
	if err != nil {
		return "", err
	}
	// 载入缓存
	_, err = c.Set(key, content, sc.Lift())
	if err != nil {
		return "", err
	}
	return content, nil
}

func (sc *SCache) Set(value string, data interface{}, expiration time.Duration) (string, error) {
	c, err := cache()
	if err != nil {
		return "", err
	}
	key := sc.Key(value)
	return c.Set(key, data, expiration)
}

func (sc *SCache) Remove(value string) (int64, error) {
	c, err := cache()
	if err != nil {
		return 0, err
	}
	key := sc.Key(value)
	return c.Remove(key)
}

