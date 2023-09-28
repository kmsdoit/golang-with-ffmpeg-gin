package service

import (
	"golang-with-ffmpeg-fiber/app/camera/model"
)

type cameraService struct {
	cameraRepository model.CameraRepository
}

func NewCameraService(cameraRepo model.CameraRepository) model.CameraService {
	return &cameraService{
		cameraRepository: cameraRepo,
	}
}

func (c *cameraService) Save(camera *model.Camera) (*model.Camera, error) {
	return c.cameraRepository.Save(camera)
}

func (c *cameraService) GetCameraById(id int) (*model.Camera, error) {
	return c.cameraRepository.GetCameraById(id)
}

func (c *cameraService) GetAllCamera() (*[]model.Camera, error) {
	return c.cameraRepository.GetAllCamera()
}
