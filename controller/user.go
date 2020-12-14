package controller

import (
	"first-auth/db"
	"first-auth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (Controller) Insert(c *gin.Context)  {
	user := models.User{
		Name: c.PostForm("name"),
		Password: c.PostForm("password"),
	}
	err := db.Insert(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"created user" : user,
	})
}
