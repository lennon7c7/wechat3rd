package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var Cache *redis.Client

var ctx = context.Background()

func NewCache(addr string, password string, db int) *redis.Client {
	log.Println("初始化缓存：redis")
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return Cache
}

// CacheSetString 设置缓存
// expiration 0表示没有时间限制。如:10 * time.Second，表示10秒
func CacheSetString(key string, value string, expiration time.Duration) (err error) {
	err = Cache.Set(ctx, key, value, expiration).Err()
	return
}

// CacheGetString 获取缓存
func CacheGetString(key string) (value string, err error) {
	value, err = Cache.Get(ctx, key).Result()
	return
}
