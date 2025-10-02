package api

import (
	"go-study-blog/common"
	"go-study-blog/config"
	"go-study-blog/models"
	"go-study-blog/services"
	"go-study-blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// Define user-related HTTP handler methods here
	userservice *services.UserService
}

func NewUserCtrl(userService *services.UserService) *UserController {
	return &UserController{userservice: userService}
}

// Implement HTTP handler methods for user operations (e.g., GetUser, CreateUser, etc.)
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.ErrInvalidInput)
		return
	}

	user, err := ctrl.userservice.GetUserByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userservice.GetAllUsers()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(common.ErrInvalidInput)
		return
	}

	if err := ctrl.userservice.UpdateUser(user); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.ErrInvalidInput)
		return
	}

	if err := ctrl.userservice.DeleteUser(uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

func (ctrl *UserController) Register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(common.ErrInvalidInput)
		return
	}

	if err := ctrl.userservice.Register(user); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (ctrl *UserController) Login(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(common.ErrInvalidInput)
		return
	}

	existingUser, err := ctrl.userservice.GetUserByName(user.Username)
	if err != nil {
		c.Error(common.ErrInvalidPassword)
		return
	}

	err = existingUser.CheckPassword(user.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	//gen token
	cfg := config.Load()
	token, err := utils.GenerateToken(existingUser.ID, cfg.App.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
