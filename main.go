package main

import (
	"home/work/controllers"
	"home/work/middleware"
	"home/work/models"

	"github.com/gin-gonic/gin"
)


func main() {
	models.InitDB()
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 :=  api.Group("v1")
		{
			v1Post := v1.Group("/post")
			v1Post.Use(middleware.JwtAuthMiddleware())
			{
				v1Post.POST("", controllers.CreateOnePost)
				v1Post.PATCH("/:id", controllers.UpdateOnePost)
				v1Post.DELETE("/:id", controllers.DeleteOneBook)
			}
			v1.GET("/post", controllers.FindManyPosts)
			v1.GET("/post/:id", controllers.FindOnePost)


			v1User := v1.Group("/user")
			v1User.Use(middleware.JwtAuthMiddleware())
			{
				v1User.PATCH("", controllers.UpdateCurrentUser)
				v1User.DELETE("", controllers.DeleteCurrentUser)
			}
			v1.GET("/user", controllers.FindManyUsers)
			v1.GET("/user/:username", controllers.FindUserByUsername)
			// v1User.GET("", controllers.FindManyUsers)

			v1Auth := v1.Group("/auth")
			{
				v1Auth.POST("/register", controllers.RegisterUser)
				v1Auth.POST("/login", controllers.LoginUser)
			}
		}
	}

	r.Run(":3000")
}
