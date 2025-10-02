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

	err := ctx.ShouldBindJSON(&post)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId, _ := ctx.Get("user_id")
	post.UserID = userId.(uint)

	err = ctrl.services.CreatePost(post, post.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (ctrl *PostController) FindList(ctx *gin.Context) {

	page := common.Pagination{}
	err := ctx.ShouldBindQuery(&page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request page"})
		return
	}

	post := models.Post{}
	err = ctx.ShouldBindQuery(&post)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, _ := ctrl.services.GetPost(post, page)

	ctx.JSON(http.StatusOK, result)

}

func (ctrl *PostController) FindByID(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	post, err := ctrl.services.GetPostByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "eror")
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (ctrl *PostController) UpdatePost(ctx *gin.Context) {

	paramPost := models.Post{}

	inputErr := ctx.ShouldBindJSON(&paramPost)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId, _ := ctx.Get("user_id")

	serviceErr := ctrl.services.UpdatePost(paramPost, userId.(uint))

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, paramPost)
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
