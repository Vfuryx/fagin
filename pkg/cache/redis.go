package cache

import (
	"context"
	"fagin/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	client *redis.Client
	*Option
}

type Option struct {
	open     bool   // 是否开启缓存
	Prefix   string // 前缀
	Addr     string // 地址
	Password string // 密码
}

func (o *Option) isOpen() bool {
	return o.open
}

var _ iCache = &redisCache{}

var Redis *redisCache = New()

var ctx = context.Background()

func New(options ...*Option) *redisCache {
	option := new(Option)
	if len(options) == 0 {
		option.Prefix = config.Redis.Prefix
		option.Addr = config.Redis.Addr
		option.open = config.Redis.Open
		option.Password = config.Redis.Password
	} else {
		option = options[0]
	}

	if !option.isOpen() {
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
		Option: option,
	}
}

// Exists 是键是否存在
func (cache *redisCache) Exists(key string) (v int64, err error) {
	v, err = cache.client.Exists(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	return v, nil
}

// Get 根据键获取数据
func (cache *redisCache) Get(key string) (str string, err error) {
	str, err = cache.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return str, nil
}

// Set 设置数据
func (cache *redisCache) Set(key string, data interface{}, expiration time.Duration) (str string, err error) {
	str, err = cache.client.Set(ctx, key, data, expiration).Result()
	if err != nil {
		if err == redis.Nil {
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
		if err == redis.Nil {
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
