package service

import (
	"golang-with-ffmpeg-fiber/app/camera/model"
	"sync"
)

var once sync.Once

type cameraService struct {
	cameraRepository model.CameraRepository
}

var instance *cameraService

func NewCameraService(cameraRepo model.CameraRepository) model.CameraService {
	once.Do(func() {
		instance = &cameraService{
			cameraRepository: cameraRepo,
		}
	})

	return instance
}

func (c *cameraService) SaveService(camera *model.Camera) (*model.Camera, error) {
	return c.cameraRepository.Save(camera)
}
