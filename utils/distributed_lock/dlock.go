package distributed_lock

import (
	"context"
	"errors"
	"fagin/config"
	"fagin/pkg/cache"
	"fagin/pkg/errorw"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"go.uber.org/atomic"

	"github.com/go-redis/redis/v8"
)

type DLock struct {
	expire      time.Duration // 失效时间
	timeWait    time.Duration // 等待基础时间
	redisClient *redis.Client
	ctx         context.Context
	wait        atomic.Int64 // 等待数量
	lock        sync.Mutex
	lock2       sync.Mutex
}

const PrefixKey = "distributed:lock:%s"

// 并发不同，比率不同
const ratio int64 = 1024

var randNum = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // 随机数生成器

var ErrConfig = errors.New("redis 配置错误")
var ErrInit = errors.New("redis 实例化错误")

// NewDLock 实例化分布式锁
func NewDLock(expire time.Duration) (*DLock, error) {
	var (
		ok                bool
		address, password string
		c                 map[string]string
	)

	if c, ok = config.Cache().Stores()[cache.DriverRedis]; !ok {
		return nil, ErrConfig
	}

	if address, ok = c["address"]; !ok {
		return nil, ErrConfig
	}

	if password, ok = c["password"]; !ok {
		return nil, ErrConfig
	}

	rdb := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     address,
		Password: password,
		DB:       0,
	})

	var ctx = context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, errorw.Wrap(ErrInit, err)
	}

	return &DLock{
			redisClient: rdb,
			expire:      expire,
			ctx:         ctx,
			timeWait:    time.Microsecond,
		},
		nil
}

func (dl *DLock) Key(key string) string {
	return fmt.Sprintf(PrefixKey, key)
}

// TimeWait 计算等待时间
func (dl *DLock) TimeWait(count int64) time.Duration {
	return dl.timeWait + time.Duration(dl.wait.Load()/count*randNum.Int63n(ratio))
}

// TryLock 尝试获取锁
// key 键
// uuid UUID
// timeout 超时时间
func (dl *DLock) TryLock(key, uuid string, timeout time.Duration) bool {
	dl.wait.Inc()

	// 防止并发访问 redis
	dl.lock.Lock()
	b, err := dl.redisClient.SetNX(dl.ctx, dl.Key(key), uuid, dl.expire).Result()
	dl.lock.Unlock()

	if err != nil {
		return false
	}

	if b {
		return true
	}

	var i int64 = 0

	out := time.Now().Add(timeout).UnixMilli()
	for time.Now().UnixMilli() < out {
		i++
		time.Sleep(dl.TimeWait(i))

		// 防止并发访问 redis
		dl.lock2.Lock()
		b, err = dl.redisClient.SetNX(dl.ctx, dl.Key(key), uuid, dl.expire).Result()
		dl.lock2.Unlock()

		if err != nil {
			return false
		}

		if b {
			return true
		}
	}

	return false
}

// UnLock 解锁
func (dl *DLock) UnLock(key, uuid string) {
	// 判断锁是自己的，才释放
	script := `
if redis.call("GET",KEYS[1]) == ARGV[1]
then
    return redis.call("DEL",KEYS[1])
else
    return 0
end
`

	dl.wait.Dec()
	_, _ = dl.redisClient.Eval(dl.ctx, script, []string{dl.Key(key)}, uuid).Result()
}
