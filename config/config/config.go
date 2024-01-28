package config

import (
	"github.com/spf13/viper"
	"log"
)

type server struct {
	Port string `mapstructure:"port"`
}
type db struct {
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
	DBName   string `mapstructure:"db_name"`
}
type jwt struct {
	Secret  string `mapstructure:"secret"`
	Expires uint   `mapstructure:"expires"`
	Issuer  string `mapstructure:"issuer"`
}
type redis struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

type config struct {
	Server server `mapstructure:"server"`
	DB     db     `mapstructure:"db"`
	Jwt    jwt    `mapstructure:"jwt"`
	Redis  redis  `mapstructure:"redis"`
}

var Config config

func Init() {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.WatchConfig()
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not find", err)
		return
	}
	err = config.Unmarshal(&Config)
	if err != nil {
		log.Fatal("Config error", err)
		return
	}
}
