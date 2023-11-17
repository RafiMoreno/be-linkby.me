package main

import (
	"github.com/RafiMoreno/be-linkby.me/controllers"
	"github.com/RafiMoreno/be-linkby.me/initializers"
	"github.com/RafiMoreno/be-linkby.me/middleware"

	docs "github.com/RafiMoreno/be-linkby.me/docs"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/sign-up", controllers.SignUp)
		v1.POST("/login", controllers.Login)
		v1.GET("/validate", middleware.RequireAuth, controllers.Validate)
		v1.GET("/profile/:username", controllers.GetProfile)
		v1.PUT("/profile/:username", middleware.RequireAuth, controllers.EditProfile)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
