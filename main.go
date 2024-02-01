package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image-host/app/midwares"
	"image-host/config/config"
	"image-host/config/database"
	"image-host/config/router"
	"log"
)

func main() {
	config.Init()
	database.Init()
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization", "Content-Type")
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run(":" + config.Config.Server.Port)
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
