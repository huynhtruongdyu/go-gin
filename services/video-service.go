package services

import "github.com/huynhtruongdyu/go-gin/models"

type VideoService interface {
	Save(models.Video) models.Video
	FindAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

// FindAll implements VideoService
func (s *videoService) FindAll() []models.Video {
	return s.videos
}

// Save implements VideoService
func (s *videoService) Save(video models.Video) models.Video {
	s.videos = append(s.videos, video)
	return video
}

func New() VideoService {
	return &videoService{}
}
