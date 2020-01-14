package cache

import (
	"fagin/config"
	"github.com/go-redis/redis/v7"
	"time"
)

type Cache struct {
	*redis.Client
	*Option
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

var _ ICache = &Cache{}

var Redis *Cache = New()

func New(options ...*Option) *Cache {
	option := new(Option)
	if len(options) == 0 {
		option.Prefix = config.Redis.Prefix
		option.Addr = config.Redis.Addr
		option.Open = config.Redis.Open
		option.Password = config.Redis.Password
	} else {
		option = options[0]
	}

	if ! option.IsOpen() {
		return nil
	}

	op := &redis.Options{
		Network:  "tcp",
		Addr:     option.Addr,
		Password: option.Password,
	}

	r := redis.NewClient(op)

	_, err := r.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &Cache{
		Client: r,
		Option: option,
	}
}

func (cache *Cache) Exists(key string) (int64, error) {
	return cache.Client.Exists(key).Result()
}

func (cache *Cache) Get(key string) (string, error) {
	return cache.Client.Get(key).Result()
}

func (cache *Cache) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return cache.Client.Set(key, value, expiration).Result()
}

func (cache *Cache) Remove(key string) (int64, error) {
	return cache.Client.Del(key).Result()
}
