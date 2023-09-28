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

func (cameraRepo *cameraRepository) GetCameraById(id int) (*model.Camera, error) {
	var camera model.Camera

	err := cameraRepo.DB.Where("id = ?", id).First(&camera).Error

	if err != nil {
		log.Println("해당하는 id의 카메라를 찾을 수 없습니다")
		return nil, fmt.Errorf(err.Error())
	}

	return &camera, nil
}

func (cameraRepo *cameraRepository) GetAllCamera() (*[]model.Camera, error) {

	var camera []model.Camera

	err := cameraRepo.DB.Find(&camera).Error

	if err != nil {
		log.Println("모든 카메라 정보를 가져올 수 없습니다")
		return nil, fmt.Errorf(err.Error())
	}

	return &camera, nil
}
