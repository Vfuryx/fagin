package cache

import "time"

// 缓存接口
type ICache interface {
	Exists(key string) (int64, error)
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	Remove(key string) (int64, error)
}
