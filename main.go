package main

import (
	"go-study-blog/api"
	"go-study-blog/config"
	"go-study-blog/logger"
	"go-study-blog/middleware"
	"go-study-blog/models"
	"go-study-blog/repositories"
	"go-study-blog/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func main() {

	cfg := config.Load()

	gormLogger := gormlogger.New(
		logger.GetLogger(),
		gormlogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	//初始化数据库
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 初始化各层
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserCtrl(userRepo)
	userController := api.NewUserCtrl(userService)

	postRepo := repositories.NewPostRepo(db)
	postService := services.NewPostService(postRepo, userRepo)
	postController := api.NewPostCtrl(postService)

	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo, userRepo, postRepo)
	commentController := api.NewCommentCtrl(commentService)

	// 设置路由
	r := gin.New()
	// 使用自定义详细日志中间件
	r.Use(middleware.DetailedLogger())
	r.Use(gin.RecoveryWithWriter(logger.GetLogger().Writer()))

	r.Use(middleware.ErroHandler())

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

			auth.GET("/posts/:id", postController.FindByID)
			auth.GET("/posts", postController.FindList)
			auth.PUT("/posts", postController.UpdatePost)
			auth.DELETE("/posts/:id", postController.DeletByID)
			auth.POST("/posts", postController.CreatePost)

			auth.GET("/comments/:id", commentController.FindByID)
			auth.GET("/comments", commentController.FindList)
			auth.PUT("/comments", commentController.UpdateComment)
			auth.DELETE("/comments/:id", commentController.DeletByID)
			auth.POST("/comments", commentController.CreateComment)
		}
	}

	// 启动服务器
	r.Run(":8080")
}
