package main

import (
	"example/be-linkby.me/controllers"
	"example/be-linkby.me/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/sign-up", controllers.SignUp)
	r.Run()
}
