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
)

func main() {

	cfg := config.Load()

	//初始化数据库
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN))

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 初始化各层
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserCtrl(userRepo)
	userController := api.NewUserCtrl(userService)

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
		}
	}

	// 启动服务器
	r.Run(":8080")
}
