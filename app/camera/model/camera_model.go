package model

import "gorm.io/gorm"

type Camera struct {
	gorm.Model
	RtspAddress string `json:"rtsp-address"`
}

type CameraRepository interface {
	Save(camera *Camera) (*Camera, error)
	GetCameraById(id int) (*Camera, error)
	GetAllCamera() (*[]Camera, error)
}

type CameraService interface {
	Save(camera *Camera) (*Camera, error)
	GetCameraById(id int) (*Camera, error)
	GetAllCamera() (*[]Camera, error)
}
