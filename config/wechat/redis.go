package wechat

import (
	"github.com/silenceper/wechat/v2/cache"
	"image-host/config/redis"
)

func setRedis(wcCache cache.Cache) cache.Cache {

	redisOpts := &cache.RedisOpts{
		Host:        redis.RedisInfo.Host + ":" + redis.RedisInfo.Port,
		Database:    redis.RedisInfo.DB,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	}
	wcCache = cache.NewRedis(redisOpts)
	return wcCache
}
