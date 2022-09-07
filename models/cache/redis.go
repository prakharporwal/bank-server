package cache

import (
	"context"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v9"
	"github.com/prakharporwal/bank-server/services/klog"
)

const RedisAddr = "13.233.195.130:6379"

func GetCacheValue(ctx *gin.Context) {
	key := ctx.Query("key")
	klog.Debug(key)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   1,
	})

	klog.Debug("redis client", &client)

	out, err := client.Get(context.Background(), "hello").Result()
	if err != nil {
		klog.Error("out", err)
	}
	klog.Debug("read value", out)

	ctx.JSON(200, out)
}
