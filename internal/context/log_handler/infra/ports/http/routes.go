package http

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller HttpServer) {
	r.POST("/select-log-file", controller.SelectLogFile)
	r.GET("/available-log-files", controller.AvailableLogFiles)
}
