package redis

import "github.com/go-redis/redis/v8"

type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

var Client *redis.Client
var Info redisConfig

func init() {
	info := getConfig()

	Client = redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	Info = info

}
