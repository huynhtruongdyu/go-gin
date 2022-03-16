package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/huynhtruongdyu/go-gin/controllers"
	"github.com/huynhtruongdyu/go-gin/services"
)

const (
	PORT = ":8080"
)

var (
	_videoService    services.VideoService      = services.New()
	_videoController controller.VideoController = controller.New(_videoService)
)

func main() {
	setupLogOutput()
	server := gin.Default()
	// server.Use(gin.Recovery(), middlewares.Logger()) // to use this custom middleware => gin.New()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, _videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, _videoController.Save(ctx))
	})
	server.Run(PORT)
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
