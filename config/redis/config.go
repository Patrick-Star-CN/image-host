package redis

import "image-host/config/config"

func getConfig() redisConfig {
	Info := redisConfig{
		Host:     "localhost",
		Port:     "6379",
		DB:       0,
		Password: "",
	}
	if config.Config.Redis.Host != "" {
		Info.Host = config.Config.Redis.Host
	}
	if config.Config.Redis.Port != "" {
		Info.Port = config.Config.Redis.Port
	}
	if config.Config.Redis.DB != 0 {
		Info.DB = config.Config.Redis.DB
	}
	if config.Config.Redis.Pass != "" {
		Info.Password = config.Config.Redis.Pass
	}
	return Info
}
