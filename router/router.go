package router

import (
	"first-auth/controller"
	"github.com/gin-gonic/gin"
)

func Init()  {
	r := gin.Default()
	ctrl := controller.Controller{}

	r.GET("/", ctrl.Home)
	r.POST("/insert", ctrl.Insert)

	r.Run("localhost:8080")
}
