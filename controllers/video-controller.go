package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhtruongdyu/go-gin/models"
	"github.com/huynhtruongdyu/go-gin/services"
)

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) models.Video
}

type controller struct {
	service services.VideoService
}

// FindAll implements VideoController
func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

// Save implements controller
func (c *controller) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
