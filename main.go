package main

import (
	"github.com/danielalejandrorosero/jwt_gin/controllers"
	"github.com/danielalejandrorosero/jwt_gin/initialize"
	"github.com/danielalejandrorosero/jwt_gin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
	initialize.DataBase()
}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/", middleware.RequireAuth, controllers.Validate)

	r.Run(":9000")
}
