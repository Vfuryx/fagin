package cache

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"time"
)

type ristrettoCache struct {
	Cache *ristretto.Cache
}

var _ iCache = &ristrettoCache{}

type RistrettoConfig struct {
	NumCounters int64
	MaxCost     int64
	BufferItems int64
}

const DriverRistrettoCache = "ristretto"

func init() {
	engineMap[DriverRistrettoCache] = newRistrettoCache
}

func NewRistrettoCache(config RistrettoConfig) (*ristrettoCache, error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		// num of keys to track frequency, usually 10*MaxCost
		NumCounters: config.NumCounters,
		// cache size(max num of items)
		MaxCost: config.MaxCost,
		// number of keys per Get buffer
		BufferItems: config.BufferItems,
		// !important: always set true if not limiting memory
		IgnoreInternalCost: true,
	})
	if err != nil {
		return nil, err
	}
	return &ristrettoCache{Cache: c}, nil
}

func newRistrettoCache() (iCache, error) {
	var ok bool
	var err error
	var c map[string]string
	var numCounters, maxCost, bufferItems string
	var n, m, b int64

	if c, ok = config.Cache().Stores[DriverRistrettoCache]; !ok {
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

func (r *ristrettoCache) Get(key string) ([]byte, error) {
	data, ok := r.Cache.Get(key)
	if !ok {
		return nil, ErrCache
	}
	if b, ok := data.([]byte); ok {
		return b, nil
	}
	return nil, ErrCache
}

func (r *ristrettoCache) Exists(key string) (bool, error) {
	_, ok := r.Cache.GetTTL(key)
	return ok, nil
}

func (r *ristrettoCache) Set(key string, value []byte, expiration time.Duration) (string, error) {
	ok := r.Cache.SetWithTTL(key, value, 1, expiration)
	r.Cache.Wait()
	fmt.Println(ok)
	return "ok", nil
}

func (r *ristrettoCache) Remove(key string) (int64, error) {
	r.Cache.Del(key)
	return 1, nil
}

func (r *ristrettoCache) Close() error {
	r.Cache.Close()
	return nil
}
