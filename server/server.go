package server

import (
	"github.com/gin-gonic/gin"
	"golang-with-ffmpeg-fiber/app/camera/router"
	"golang-with-ffmpeg-fiber/database"
)

func Init() *gin.Engine {
	r := gin.Default()
	conn := database.DatabaseInit()
	router.CameraRouter(conn, r)

	r.Run(":18080")
	return r
}
