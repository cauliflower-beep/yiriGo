package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goburrow/cache"
)

type UserCache struct {
	cache cache.LoadingCache
}

type UserInfo struct {
	name string
	age  int
}

func main() {
	// 生成缓存
	load := func(k cache.Key) (user cache.Value, err error) {
		user = UserInfo{
			name: "lufy",
			age:  21,
		}
		//time.Sleep(100 * time.Millisecond) // Slow task
		return
	}
	// Create a loading cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(100),                 // 限制缓存中的条目数
		cache.WithExpireAfterAccess(1*time.Minute), // Expire entries after 1 minute since last accessed.
		//cache.WithRefreshAfterWrite(2*time.Minute), // Expire entries after 2 minutes since last created.
	)

	// 第一次取缓存值，如果取不到，就会调用load加载值
	v, _ := c.Get(rand.Intn(200))
	fmt.Println(v)

	// 由于设置了过期时间，所以时间到了之后应该是取不到值的
	getTicker := time.Tick(100 * time.Millisecond)
	reportTicker := time.Tick(5 * time.Second)
	for {
		select {
		case <-getTicker:
			v, _ := c.Get(rand.Intn(200))
			fmt.Println(v)
		case <-reportTicker:
			st := cache.Stats{}
			c.Stats(&st)
			fmt.Printf("%+v\n", st)
		}
	}
}
