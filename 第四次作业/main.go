package main

import (
	"Four/config"
	"Four/controllers"
	"Four/database"
	"Four/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.ConnectDB(cfg)
	if err != nil {
		panic("Failed to connect database" + err.Error())
	}
	r := gin.Default()

	authController := &controllers.AuthController{DB: db, Config: cfg}
	postController := &controllers.PostController{DB: db}
	commentController := &controllers.CommentController{DB: db}

	public := r.Group("/api")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
		public.GET("/posts", postController.GetPosts)
		public.GET("/posts/:postid", postController.GetPost)
		public.GET("/posts/:postid/comments", commentController.GetComments)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuth(cfg))
	{
		protected.POST("/posts", postController.CreatePost)
		protected.PUT("/posts/:postid", postController.UpdatePost)
		protected.DELETE("/posts/:postid", postController.DeletePost)
		protected.POST("/posts/:postid/comments", commentController.CreateComment)
	}
	r.Run(":8080")
}
