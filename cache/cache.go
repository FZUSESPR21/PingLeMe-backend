package cache

import (
	"context"
	"strings"
	"time"
)

var ctx = context.Background()

func Get(keyStrs ...string) (string, error) {
	var builder strings.Builder
	for _, keyStr := range keyStrs {
		builder.WriteString(keyStr)
	}

	return RedisClient.Get(ctx, builder.String()).Result()
}

func SetNX(key string, value interface{}, duration time.Duration) (bool, error) {
	return RedisClient.SetNX(ctx, key, value, duration).Result()
}

func Del(key string) (int64, error) {
	return RedisClient.Del(ctx, key).Result()
}
