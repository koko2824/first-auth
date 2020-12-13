package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {}

func (Controller) Home(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Home",
	})
}