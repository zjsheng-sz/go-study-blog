package api

import (
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	services *services.PostService
}

func NewPostCtrl(services *services.PostService) *PostController {
	return &PostController{
		services: services,
	}
}

func (ctrl *PostController) CreatePost(ctx *gin.Context) {
	post := models.Post{}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	userId, _ := ctx.Get("user_id")
	post.UserID = userId.(uint)

	if err := ctrl.services.CreatePost(post, post.UserID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    post,
	})
}

func (ctrl *PostController) FindList(ctx *gin.Context) {
	page := common.Pagination{}
	if err := ctx.ShouldBindQuery(&page); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	post := models.Post{}
	if err := ctx.ShouldBindQuery(&post); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	result, err := ctrl.services.GetPost(post, page)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func (ctrl *PostController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	post, err := ctrl.services.GetPostByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    post,
	})
}

func (ctrl *PostController) UpdatePost(ctx *gin.Context) {
	paramPost := models.Post{}
	if err := ctx.ShouldBindJSON(&paramPost); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	userId, _ := ctx.Get("user_id")
	if err := ctrl.services.UpdatePost(paramPost, userId.(uint)); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    paramPost,
	})
}

func (ctrl *PostController) DeletByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId, _ := ctx.Get("user_id")

	serviceErr := ctrl.services.DeletePost(id, userId.(uint))

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
