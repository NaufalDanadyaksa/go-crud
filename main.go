package main

import (
	"go-crud/controllers"
	"go-crud/inits"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.ConnectDB()
}
func main() {
	router := gin.Default()
	router.GET("/ping", controllers.Users)
	router.POST("/user", controllers.UserCreate)
	router.PUT("/user/:id", controllers.UserUpdate)
	router.GET("/users", controllers.UserIndex)
	router.GET("/user/:id", controllers.UserShow)
	router.DELETE("/user/:id", controllers.UserDelete)
	
	router.Run()
}
