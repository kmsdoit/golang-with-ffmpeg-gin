package router

import (
	"github.com/gin-gonic/gin"
	"golang-with-ffmpeg-fiber/app/camera/controller"
	"golang-with-ffmpeg-fiber/app/camera/repository"
	"golang-with-ffmpeg-fiber/app/camera/service"
	"gorm.io/gorm"
)

func CameraRouter(conn *gorm.DB, gin *gin.Engine) {

	cameraRepo := repository.NewCameraRepository(conn)
	cameraService := service.NewCameraService(cameraRepo)
	cameraController := controller.NewCameraController(cameraService)

	cameraRouterV1 := gin.Group("/api/v1/camera")
	{
		cameraRouterV1.POST("/create", cameraController.CreateCameraController)
		cameraRouterV1.GET("/get/camera", cameraController.GetCameraByIdController)
		cameraRouterV1.GET("/cameras", cameraController.GetAllCameraController)
	}
}
