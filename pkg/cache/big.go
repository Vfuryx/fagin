package cache

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fagin/config"
	"time"

	"github.com/allegro/bigcache/v3"
)

type BigCache struct {
	Cache *bigcache.BigCache
}

var _ cache = &BigCache{}

type BigConfig struct {
	Eviction time.Duration // 失效时间
}

type Big struct {
	Data       []byte
	Expiration time.Time
}

const DriverBigCache = "big"

func init() {
	engineMap[DriverBigCache] = newBigCache
}

func NewBigCache(conf BigConfig) (*BigCache, error) {
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(conf.Eviction))
	if err != nil {
		return nil, err
	}

	return &BigCache{Cache: c}, nil
}

func newBigCache() (cache, error) {
	var (
		ok       bool
		eviction string
		c        map[string]string
		e        time.Duration
		err      error
	)

	if c, ok = config.Cache().Stores()[DriverBigCache]; !ok {
		return nil, ErrConfig
	}

	if eviction, ok = c["eviction"]; !ok {
		return nil, ErrConfig
	}

	if e, err = time.ParseDuration(eviction + "s"); err != nil {
		return nil, ErrConfig
	}

	return NewBigCache(BigConfig{
		Eviction: e,
	})
}

func (b *BigCache) Exists(key string) (bool, error) {
	bs, err := b.Cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return false, nil
		}

		return false, err
	}

	var body Big

	dec := gob.NewDecoder(bytes.NewBuffer(bs))

	if err := dec.Decode(&body); err != nil {
		return false, err
	}

	if body.Expiration.After(time.Now()) {
		return true, nil
	}

	return false, nil
}

func (b *BigCache) Get(key string) ([]byte, error) {
	bs, err := b.Cache.Get(key)
	if err != nil {
		return nil, err
	}

	var body Big

	dec := gob.NewDecoder(bytes.NewBuffer(bs))

	err = dec.Decode(&body)
	if err != nil {
		return nil, err
	}

	if body.Expiration.Before(time.Now()) {
		return nil, bigcache.ErrEntryNotFound
	}

	return body.Data, nil
}

func (b *BigCache) Set(key string, value []byte, expiration time.Duration) (string, error) {
	big := Big{
		Data:       value,
		Expiration: time.Now().Add(expiration),
	}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err := enc.Encode(big)
	if err != nil {
		return "", err
	}

	return "ok", b.Cache.Set(key, buf.Bytes())
}

func (b *BigCache) Remove(key string) (int64, error) {
	err := b.Cache.Delete(key)
	if err != nil && err != bigcache.ErrEntryNotFound {
		return 0, err
	}

	return 1, nil
}

func (b *BigCache) Close() error {
	return nil
}
