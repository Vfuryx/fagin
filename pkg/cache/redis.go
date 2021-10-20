package cache

import (
	"context"
	"errors"
	"fagin/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	client *redis.Client
}

type RedisOption struct {
	open     bool   // 是否开启缓存
	Prefix   string // 前缀
	Addr     string // 地址
	Password string // 密码
}

const DriverRedis = "redis"

var _ cache = &redisCache{}

func init() {
	engineMap[DriverRedis] = newRedis
}

var ctx = context.Background()

func NewRedis(option RedisOption) (*redisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     option.Addr,
		Password: option.Password,
		DB:       0,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, ErrNotOpen
	}

	return &redisCache{client: rdb}, nil
}

func newRedis() (cache, error) {
	var ok bool
	var address, password string
	var c map[string]string

	if c, ok = config.Cache().Stores[DriverRedis]; !ok {
		return nil, ErrConfig
	}
	if address, ok = c["address"]; !ok {
		return nil, ErrConfig
	}
	if password, ok = c["password"]; !ok {
		return nil, ErrConfig
	}

	return NewRedis(RedisOption{
		Addr:     address,
		Password: password,
	})
}

// Exists 是键是否存在
func (cache *redisCache) Exists(key string) (b bool, err error) {
	_, err = cache.client.Exists(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Get 根据键获取数据
func (cache *redisCache) Get(key string) ([]byte, error) {
	str, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

// Set 设置数据
func (cache *redisCache) Set(key string, data []byte, expiration time.Duration) (str string, err error) {
	str, err = cache.client.Set(ctx, key, data, expiration).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return str, nil
}

// Remove 删除数据
func (cache *redisCache) Remove(key string) (v int64, err error) {
	v, err = cache.client.Del(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}
	return v, nil
}

// Close 关闭
func (cache *redisCache) Close() error {
	return cache.client.Close()
}
