package database

import (
	"golang-with-ffmpeg-fiber/app/camera/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("DB 접속에 실패했습니다")
	}

	db.AutoMigrate(&model.Camera{})

	return db
}
