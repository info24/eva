package router

import (
	"github.com/gin-gonic/gin"
	"github.com/info24/eva/controller"
)

func InitRouter(router *gin.RouterGroup) {
	//router.GET("/ws/:id", controller.RegisterSsh)
	router.GET("/device", controller.GetAllDevice)
	router.POST("/device", controller.AddDevice)
	router.POST("/device/:id", controller.UpdateDevice)
	router.DELETE("/device/:id", controller.DeleteDevice)
	//router.GET("/ws/:id", controller.RegisterSsh)
}
