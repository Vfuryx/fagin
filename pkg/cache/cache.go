package cache

import (
	"github.com/pkg/errors"
	"time"
)

// 缓存接口
type oCache interface {
	Lift() time.Duration
	Key(value string) string
}

type iCache interface {
	isOpen() bool // 缓存是否开启
	Get(key string) (string, error)
	Exists(key string) (int64, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	Remove(key string) (int64, error)
}

var NotOpenErr = errors.New("缓存服务未开启")

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

func cache() (iCache, error) {
	if Redis == nil {
		return nil, NotOpenErr
	}
	return Redis, nil
}

type SCache struct {
	Prefix   string
	LifeTime time.Duration
	Content  GetterFunc // 获取回源数据方法
	oCache
}

var _ iCache = &SCache{}

// SetFunc 载入多态
func (sc *SCache) SetFunc(o oCache) {
	sc.oCache = o
}

// Exists 是键是否存在
func (sc *SCache) Exists(value string) (int64, error) {
	c, err := cache()
	if err != nil {
		return 0, err
	}
	key := sc.Key(value)
	return c.Exists(key)
}

func (sc *SCache) Set(value string, data interface{}, expiration time.Duration) (string, error) {
	c, err := cache()
	if err != nil {
		return "", err
	}
	key := sc.Key(value)
	return c.Set(key, data, expiration)
}

func (sc *SCache) Remove(key string) (int64, error) {
	c, err := cache()
	if err != nil {
		return 0, err
	}
	key = sc.Key(key)
	return c.Remove(key)
}

// Get 根据键获取数据
func (sc *SCache) Get(value string) (string, error) {
	c, err := cache()
	if err != nil {
		// 缓存关闭直接获取数据
		if err == NotOpenErr {
			content, err := sc.Content.Get(value)
			if err != nil {
				return "", err
			}
			return string(content), nil
		}
		return "", err
	}
	// 获取缓存key
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
	content, err := sc.Content.Get(value)
	if err != nil {
		return "", err
	}
	// 载入缓存
	_, err = c.Set(key, content, sc.Lift())
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 是否开启
func (sc *SCache) isOpen() bool {
	return false
}
