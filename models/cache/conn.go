package cache

import (
	"fmt"
	"sync"

	redis "github.com/go-redis/redis/v9"
)

var lock = &sync.Mutex{}

var redisInstance *redis.Client

const RedisAddr = "localhost:6379"

func getInstance() *redis.Client {
	if redisInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if redisInstance == nil {
			fmt.Println("Creating single instance now.")
			redisInstance = redis.NewClient(&redis.Options{
				Addr: RedisAddr,
				DB:   1,
			})
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return redisInstance
}

func GetInstance() *redis.Client {
	return getInstance()
}
