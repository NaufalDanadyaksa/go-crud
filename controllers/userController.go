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
		Name    string
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
		"result":  user,
	})
}

func UserIndex(c *gin.Context) {
	var users []models.User

	inits.DB.Find(&users)

	c.JSON(200, gin.H{
		"message": "success",
		"result":  users,
	})
}

func UserShow(c *gin.Context) {
	var user models.User

	inits.DB.First(&user, c.Param("id"))

	if user.ID == 0 {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"result":  user,
		})
	}

}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Name    string
		Address string
	}

	c.Bind(&request)

	var user models.User

	inits.DB.First(&user, id)

	inits.DB.Model(&user).Updates(models.User{
		Name:    request.Name,
		Address: request.Address,
	})

	c.JSON(200, gin.H{
		"message": "success",
		"result":  user,
	})

}

func UserDelete(c *gin.Context) {
	var user models.User
	inits.DB.Delete(&user, c.Param("id"))

	c.JSON(404, gin.H{
		"message": "user not found",
	})

}
