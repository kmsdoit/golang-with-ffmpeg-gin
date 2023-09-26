package model

import "gorm.io/gorm"

type Camera struct {
	gorm.Model
	RtspAddress string `json:"rtsp-address"`
}

type CameraRepository interface {
	Save(camera *Camera) (*Camera, error)
}

type CameraService interface {
	SaveService(camera *Camera) (*Camera, error)
}
