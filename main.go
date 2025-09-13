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
	router.Run()
}
