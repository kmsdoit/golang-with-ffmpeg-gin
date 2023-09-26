package repository

import (
	"fmt"
	"golang-with-ffmpeg-fiber/app/camera/model"
	"gorm.io/gorm"
	"log"
)

type cameraRepository struct {
	DB *gorm.DB
}

func NewCameraRepository(db *gorm.DB) model.CameraRepository {
	return &cameraRepository{
		DB: db,
	}
}

func (cameraRepo *cameraRepository) Save(camera *model.Camera) (*model.Camera, error) {
	err := cameraRepo.DB.Create(camera).Error

	if err != nil {
		log.Println("카메라 생성이 실패했습니다")
		return nil, fmt.Errorf(err.Error())
	}

	return camera, nil
}
