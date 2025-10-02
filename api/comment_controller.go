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
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	if err := ctrl.services.CreateComment(comment); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comment,
	})
}

func (ctrl *CommentController) FindList(ctx *gin.Context) {
	page := common.Pagination{}
	if err := ctx.ShouldBindQuery(&page); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	comment := models.Comment{}
	if err := ctx.ShouldBindQuery(&comment); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	result, err := ctrl.services.GetComment(comment, page)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

func (ctrl *CommentController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	comment, err := ctrl.services.GetCommentByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comment,
	})
}

func (ctrl *CommentController) UpdateComment(ctx *gin.Context) {
	paramComment := models.Comment{}
	if err := ctx.ShouldBindJSON(&paramComment); err != nil {
		ctx.Error(common.ErrInvalidInput)
		return
	}

	userId, _ := ctx.Get("user_id")
	if err := ctrl.services.UpdateComment(paramComment, userId.(uint)); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    paramComment,
	})
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
