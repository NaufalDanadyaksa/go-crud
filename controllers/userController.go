package controllers

import (
	"go-crud/inits"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserCreate(c *gin.Context) {
	var request struct {
		Name string
		Address string
	}

	c.Bind(&request)


	user := models.User{Name: request.Name, Address: request.Address}

	result := inits.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "failed to create",
		})
		
	}
	c.JSON(200, gin.H{
		"message": "success",
		"result": user,
	})
}
