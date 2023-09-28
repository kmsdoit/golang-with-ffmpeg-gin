package controller

import (
	"github.com/gin-gonic/gin"
	"golang-with-ffmpeg-fiber/app/camera/model"
	"golang-with-ffmpeg-fiber/app/common"
	"net/http"
	"strconv"
)

type cameraController struct {
	cameraService model.CameraService
}

type CameraController interface {
	CreateCameraController(c *gin.Context)
	GetCameraByIdController(c *gin.Context)
	GetAllCameraController(c *gin.Context)
}

func NewCameraController(c model.CameraService) CameraController {
	return &cameraController{
		cameraService: c,
	}
}

func (cc *cameraController) GetAllCameraController(c *gin.Context) {

	cameras, err := cc.cameraService.GetAllCamera()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			common.FailResponse{
				Response: common.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "DB를 확인해주세요",
				},
				Error: err,
			})
	}

	c.JSON(http.StatusOK,
		common.SuccessResponse{
			Response: common.Response{
				StatusCode: http.StatusOK,
				Message:    "모든 카메라 정보입니다",
			},
			Data: cameras,
		},
	)

}

func (cc *cameraController) GetCameraByIdController(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("id"))

	camera, err := cc.cameraService.GetCameraById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			common.FailResponse{
				Response: common.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
				},
				Error: err,
			})

		return
	}

	c.JSON(http.StatusOK,
		common.SuccessResponse{
			Response: common.Response{
				StatusCode: http.StatusOK,
				Message:    "해당하는 id에 대한 카메라 정보 입니다",
			},
			Data: camera,
		},
	)
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

	cc.cameraService.Save(&camera)

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
