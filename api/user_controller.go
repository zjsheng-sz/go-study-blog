package api

import (
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

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.userservice.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)

}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userservice.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(200, users)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {

	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := ctrl.userservice.UpdateUser(user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(200, user)

}

func (ctrl *UserController) DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.userservice.DeleteUser(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})

}

func (ctrl *UserController) Register(c *gin.Context) {

	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := ctrl.userservice.Register(user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, user)

}

func (ctrl *UserController) Login(c *gin.Context) {

	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	existingUser, err := ctrl.userservice.GetUserByName(user.Username)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
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

func (ctrl *UserController) IdentiferAuth(ctx *gin.Context) {

	user := models.User{}
	userId, _ := ctx.Get("user_id")
	user.ID = userId.(uint)
	user.IsAuthen = true

	err := ctrl.userservice.UpdateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, "success")

}
