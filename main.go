package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image-host/app/midwares"
	"image-host/config/database"
	"image-host/config/router"
	"image-host/config/session"
	"log"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	session.Init(r)
	router.Init(r)

	err := r.Run(":8088")
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
