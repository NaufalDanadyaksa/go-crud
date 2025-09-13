package main

import (
	"go-crud/inits"
	"go-crud/models"
)

func init() {
	inits.LoadEnv()
	inits.ConnectDB()
}

func main() {
	inits.DB.AutoMigrate(&models.User{})
}