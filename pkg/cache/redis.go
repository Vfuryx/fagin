package cache

import (
	"context"
	"fagin/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	client *redis.Client
	option *Option
}

type Option struct {
	Open     bool   // 是否开启缓存
	Prefix   string // 前缀
	Addr     string // 地址
	Password string // 密码
}

func (o Option) IsOpen() bool {
	return o.Open
}

var _ iCache = &redisCache{}

var Redis *redisCache = New()

var ctx = context.Background()

func New(options ...*Option) *redisCache {
	option := new(Option)
	if len(options) == 0 {
		option.Prefix = config.Redis.Prefix
		option.Addr = config.Redis.Addr
		option.Open = config.Redis.Open
		option.Password = config.Redis.Password
	} else {
		option = options[0]
	}

	if !option.IsOpen() {
		return nil
	}

	op := &redis.Options{
		Network:  "tcp",
		Addr:     option.Addr,
		Password: option.Password,
		DB:       0,
	}

	rdb := redis.NewClient(op)

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return &redisCache{
		client: rdb,
		option: option,
	}
}

// 是键是否存在
func (cache *redisCache) Exists(key string) (int64, error) {
	return cache.client.Exists(ctx, key).Result()
}

// 根据键获取数据
func (cache *redisCache) Get(key string) (string, error) {
	return cache.client.Get(ctx, key).Result()
}

// 设置数据
func (cache *redisCache) Set(key string, data interface{}, expiration time.Duration) (string, error) {
	return cache.client.Set(ctx, key, data, expiration).Result()
}

// 删除数据
func (cache *redisCache) Remove(key string) (int64, error) {
	return cache.client.Del(ctx, key).Result()
}
// 关闭
func (cache *redisCache) Close() error {
	return cache.client.Close()
}
