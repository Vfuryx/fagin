package cache

import (
	"fagin/config"
	"fagin/utils"
	"time"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache struct {
	Cache *ristretto.Cache
}

var _ cache = &RistrettoCache{}

// RistrettoConfig 配置
type RistrettoConfig struct {
	NumCounters int64
	MaxCost     int64
	BufferItems int64
}

// DriverRistrettoCache ristretto
const DriverRistrettoCache = "ristretto"

func init() {
	engineMap[DriverRistrettoCache] = newRistrettoCache
}

// NewRistrettoCache 实例化
func NewRistrettoCache(conf RistrettoConfig) (*RistrettoCache, error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		// num of keys to track frequency, usually 10*MaxCost
		NumCounters: conf.NumCounters,
		// cache size(max num of items)
		MaxCost: conf.MaxCost,
		// number of keys per Get buffer
		BufferItems: conf.BufferItems,
		// !important: always set true if not limiting memory
		IgnoreInternalCost: true,
	})
	if err != nil {
		return nil, err
	}

	return &RistrettoCache{Cache: c}, nil
}

func newRistrettoCache() (cache, error) {
	var (
		ok                                bool
		err                               error
		c                                 map[string]string
		numCounters, maxCost, bufferItems string
		n, m, b                           int64
	)

	if c, ok = config.Cache().Stores()[DriverRistrettoCache]; !ok {
		return nil, ErrConfig
	}

	if numCounters, ok = c["num_counters"]; !ok {
		return nil, ErrConfig
	}

	if maxCost, ok = c["max_cost"]; !ok {
		return nil, ErrConfig
	}

	if bufferItems, ok = c["buffer_items"]; !ok {
		return nil, ErrConfig
	}

	if n, err = utils.StrToInt64(numCounters); err != nil {
		return nil, err
	}

	if m, err = utils.StrToInt64(maxCost); err != nil {
		return nil, err
	}

	if b, err = utils.StrToInt64(bufferItems); err != nil {
		return nil, err
	}

	return NewRistrettoCache(RistrettoConfig{
		NumCounters: n,
		MaxCost:     m,
		BufferItems: b,
	})
}

func (r *RistrettoCache) Get(key string) ([]byte, error) {
	data, ok := r.Cache.Get(key)
	if !ok {
		return nil, ErrCache
	}

	if b, ok := data.([]byte); ok {
		return b, nil
	}

	return nil, ErrCache
}

func (r *RistrettoCache) Exists(key string) (bool, error) {
	_, ok := r.Cache.GetTTL(key)
	return ok, nil
}

func (r *RistrettoCache) Set(key string, value []byte, expiration time.Duration) (string, error) {
	r.Cache.SetWithTTL(key, value, 1, expiration)
	r.Cache.Wait()

	return "ok", nil
}

func (r *RistrettoCache) Remove(key string) (int64, error) {
	r.Cache.Del(key)
	return 1, nil
}

func (r *RistrettoCache) Close() error {
	r.Cache.Close()
	return nil
}
