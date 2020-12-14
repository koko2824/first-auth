package controller

import (
	"first-auth/db"
	"first-auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func passHash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, err
}

func (Controller) Insert(c *gin.Context) {
	hash, err := passHash(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}
	user := models.User{
		Name:     c.PostForm("name"),
		Password: string(hash),
	}
	
	err = db.Insert(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"created user": user,
	})
}
