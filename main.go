package main

import (
	"go-study-blog/api"
	"go-study-blog/config"
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

	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.GetAllUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// 启动服务器
	r.Run(":8080")
}
