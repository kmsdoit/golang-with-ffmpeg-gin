package controller

import (
	"github.com/gin-gonic/gin"
	"golang-with-ffmpeg-fiber/app/camera/model"
	"golang-with-ffmpeg-fiber/app/common"
	"net/http"
)

type cameraController struct {
	cameraService model.CameraService
}

type CameraController interface {
	CreateCameraController(c *gin.Context)
}

func NewCameraController(c model.CameraService) CameraController {
	return &cameraController{
		cameraService: c,
	}
}

func (cc *cameraController) CreateCameraController(c *gin.Context) {
	var camera model.Camera

	if err := c.ShouldBindJSON(&camera); err != nil {
		c.JSON(http.StatusBadRequest,
			common.FailResponse{
				Response: common.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "http.StatusBadRequest",
				},
				Error: err,
			},
		)
		return
	}

	cc.cameraService.SaveService(&camera)

	c.JSON(http.StatusOK,
		common.SuccessResponse{
			Response: common.Response{
				StatusCode: http.StatusOK,
				Message:    "카메라 생성 성공",
			},
			Data: camera,
		},
	)
}
