package main

import (
	"go-study-blog/api"
	"go-study-blog/config"
	"go-study-blog/middleware"
	"go-study-blog/models"
	"go-study-blog/repositories"
	"go-study-blog/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	cfg := config.Load()

	//初始化数据库
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有 SQL
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 初始化各层
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserCtrl(userRepo)
	userController := api.NewUserCtrl(userService)

	// 初始化各层
	postRepo := repositories.NewPostRepo(db)
	postService := services.NewPostService(postRepo, userRepo)
	postController := api.NewPostCtrl(postService)

	// 设置路由
	r := gin.Default()

	// 路由
	api := r.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)

		// 需要认证的路由
		auth := api.Group("/auth")
		auth.Use(middleware.JWTAuthMiddleware())
		{
			auth.GET("/users/:id", userController.GetUser)
			auth.GET("/users", userController.GetAllUsers)
			auth.PUT("/users", userController.UpdateUser)
			auth.DELETE("/users/:id", userController.DeleteUser)
			auth.POST("/users/identiferAuch", userController.IdentiferAuth)

			auth.GET("/posts/:id", postController.FindByID)
			auth.GET("/posts", postController.FindList)
			auth.PUT("/posts", postController.UpdatePost)
			auth.DELETE("/posts/:id", postController.DeletByID)
			auth.POST("/posts", postController.CreatePost)

		}
	}

	// 启动服务器
	r.Run(":8080")
}
