package cache

import (
	"errors"
	"fagin/config"
	"fmt"
	"sync"
	"time"
)

type cache interface {
	Get(key string) ([]byte, error)
	Exists(key string) (bool, error)
	Set(key string, value []byte, expiration time.Duration) (string, error)
	Remove(key string) (int64, error)
	Close() error
}

type GetterFunc func() ([]byte, error)

func (f GetterFunc) Get() ([]byte, error) {
	return f()
}

var ErrNotOpen = errors.New("缓存服务未开启")
var ErrConfig = errors.New("缓存配置错误")
var ErrCache = errors.New("缓存获取失败")

var engineMap = map[string]func() (cache, error){}

var engine cache

// Init 初始化
func Init() {
	if _, err := cacheEngine(); err != nil && err != ErrNotOpen {
		panic(err)
	}
}

// Close 关闭
func Close() {
	if engine != nil {
		_ = engine.Close()
	}
}

// cacheEngine 获取缓存
func cacheEngine() (cache, error) {
	if !config.Cache().Open {
		return nil, ErrNotOpen
	}
	if engine == nil {
		fn, ok := engineMap[config.Cache().DefDriver]
		if !ok {
			return nil, errors.New("err")
		}
		var err error

		engine, err = fn()
		if err != nil {
			engine = nil // 重置为 nil
		}
		return engine, err
	}
	return engine, nil
}

type Cache struct {
	format   string
	lifeTime time.Duration
	content  GetterFunc // 获取回源数据方法
}

func (sc *Cache) SetKeyFormat(format string) {
	sc.format = config.Cache().Prefix + format
}

func (sc *Cache) SetLifeTime(t time.Duration) {
	sc.lifeTime = t
}

func (sc *Cache) SetFunc(f GetterFunc) {
	sc.content = f
}

func (sc *Cache) Lift() time.Duration {
	return sc.lifeTime
}

// Key 获取键名称
func (sc *Cache) Key(value ...interface{}) string {
	return fmt.Sprintf(sc.format, value...)
}

// Exists 是键是否存在
func (sc *Cache) Exists(value ...interface{}) (bool, error) {
	c, err := cacheEngine()
	if err != nil {
		return false, err
	}
	return c.Exists(sc.Key(value...))
}

// Remove 删除
func (sc *Cache) Remove(value ...interface{}) (int64, error) {
	c, err := cacheEngine()
	if err != nil {
		return 0, err
	}
	return c.Remove(sc.Key(value...))
}

var lock sync.Mutex
var num = 0

// Get 根据键获取数据
func (sc *Cache) Get(value ...interface{}) (data []byte, err error) {
	c, err := cacheEngine()
	if err != nil {
		// 缓存关闭直接获取数据
		if errors.Is(err, ErrNotOpen) {
			if data, err = sc.content.Get(); err == nil {
				return data, nil
			}
		}
		return nil, err
	}

	var ok bool
	// 获取缓存key
	key := sc.Key(value...)

	if ok, err = c.Exists(key); err != nil { // 先处理错误
		return nil, err
	} else if ok { // 存在
		if data, err = c.Get(key); err == nil {
			return data, err
		}
	}

	lock.Lock()
	defer lock.Unlock()
	if data, err = c.Get(key); err == nil {
		return data, err
	}
	num++
	fmt.Println(num)
	// 不存在
	if data, err = sc.content.Get(); err != nil {
		return nil, err
	}
	// 载入缓存
	if _, err = c.Set(key, data, sc.Lift()); err != nil {
		return nil, err
	}
	return data, nil
}

func (sc *Cache) Close() error {
	c, err := cacheEngine()
	if err != nil {
		return err
	}
	return c.Close()
}
