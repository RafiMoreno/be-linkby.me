package main

import (
	"example/be-linkby.me/controllers"
	"example/be-linkby.me/initializers"
	"example/be-linkby.me/middleware"

	docs "example/be-linkby.me/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
	v1.POST("/sign-up", controllers.SignUp)
	v1.POST("/login", controllers.Login)
	v1.GET("/validate", middleware.RequireAuth, controllers.Validate)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
