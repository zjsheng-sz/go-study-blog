package api

import (
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	services *services.CommentService
}

func NewCommentCtrl(services *services.CommentService) *CommentController {
	return &CommentController{
		services: services,
	}
}

func (ctrl *CommentController) CreateComment(ctx *gin.Context) {

	comment := models.Comment{}

	err := ctx.ShouldBindJSON(&comment)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = ctrl.services.CreateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (ctrl *CommentController) FindList(ctx *gin.Context) {

	page := common.Pagination{}
	err := ctx.ShouldBindQuery(&page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request page"})
		return
	}

	comment := models.Comment{}
	err = ctx.ShouldBindQuery(&comment)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, _ := ctrl.services.GetComment(comment, page)

	ctx.JSON(http.StatusOK, result)

}

func (ctrl *CommentController) FindByID(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	comment, err := ctrl.services.GetCommentByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "eror")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (ctrl *CommentController) UpdateComment(ctx *gin.Context) {

	paramComment := models.Comment{}

	inputErr := ctx.ShouldBindJSON(&paramComment)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId, _ := ctx.Get("user_id")

	serviceErr := ctrl.services.UpdateComment(paramComment, userId.(uint))

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, paramComment)
}

func (ctrl *CommentController) DeletByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId, _ := ctx.Get("user_id")

	serviceErr := ctrl.services.DeleteComment(id, userId.(uint))

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
