package distributed_lock_test

import (
	"fagin/config"
	"fagin/utils/distributed_lock"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// 初始化配置
	config.Init()

	m.Run()
	os.Exit(1)
}

var wg sync.WaitGroup

func TestDLockSimple(t *testing.T) {
	dl, err := distributed_lock.NewDLock(10 * time.Second)
	if err != nil {
		return
	}

	num := 0
	now := time.Now()

	for i := 0; i < 100000; i++ {
		uid, err := uuid.NewUUID()
		if err != nil {
			return
		}

		wg.Add(1)

		go func(no int, uid string) {
			ok := dl.TryLock("test", uid, 30*time.Second)
			defer dl.UnLock("test", uid)
			wg.Done()

			if !ok {
				fmt.Printf("失败 NO:%d \n", no)
				return
			}

			num += 1
			fmt.Printf("NO:%d num:%d\n", no, num)
		}(i, uid.String())
	}

	wg.Wait()

	fmt.Println(time.Since(now))
}

func TestDLock(t *testing.T) {
	dl, err := distributed_lock.NewDLock(10 * time.Second)
	if err != nil {
		return
	}

	num := 0
	now := time.Now()

	for i := 0; i < 100000; i++ {
		uid, err := uuid.NewUUID()
		if err != nil {
			return
		}

		wg.Add(1)

		go func(no int, uid string) {
			ok := dl.TryLock("test", uid, 30*time.Second)
			defer dl.UnLock("test", uid)
			wg.Done()

			if !ok {
				fmt.Printf("失败 NO:%d \n", no)
				return
			}

			num += 1

			time.Sleep(100 * time.Microsecond)

			fmt.Printf("NO:%d num:%d\n", no, num)
		}(i, uid.String())
	}

	wg.Wait()

	fmt.Println(time.Since(now))
}

func TestDLockLock(t *testing.T) {
	dl, err := distributed_lock.NewDLock(2 * time.Second)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()

	uid, err := uuid.NewUUID()
	if err != nil {
		return
	}

	ok := dl.TryLock("test", uid.String(), 5*time.Second)

	assert.Equal(t, true, ok, time.Since(now).Seconds())
}

func TestDLockUnLock(t *testing.T) {
	dl, err := distributed_lock.NewDLock(60 * time.Second)
	if err != nil {
		t.Fatal(err)
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		return
	}

	dl.UnLock("test", uid.String())
}
