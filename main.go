package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sp-yduck/gocron-sample/controller"
)

func main() {
	router := gin.Default()

	router.GET("/create/:tag", controller.CreateJob)
	router.GET("/kill/:tag", controller.KillJob)
	router.GET("/killall", controller.KillAllJob)

	router.Run()
}
