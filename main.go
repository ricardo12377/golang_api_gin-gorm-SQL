package main

import (
	"apigin/controllers"
	"apigin/models"

	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindAllUsers)
	r.POST("/users", controllers.CreateBook)
	r.GET("/users/:id", controllers.FindUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}