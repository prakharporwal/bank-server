package cache

import (
	"context"
	"time"

	"github.com/prakharporwal/bank-server/services/klog"
)

var (
	ExpirationTimeInMinutes = time.Duration(30 * time.Minute)
)

func GetCacheValue(key string) string {
	klog.Debug(key)

	client := GetInstance()

	klog.Debug("redis client", &client)
	out, err := client.Get(context.Background(), key).Result()
	if err != nil {
		klog.Error("out", err)
	}
	klog.Debug("read value", out)

	return out
}

func SetCache(key string, value string) interface{} {

	klog.Debug(key, value)

	client := GetInstance()

	out, err := client.Set(context.Background(), key, value, ExpirationTimeInMinutes).Result()
	if err != nil {
		klog.Error("out", err)
	}
	klog.Debug("read value", out)

	return out
}
