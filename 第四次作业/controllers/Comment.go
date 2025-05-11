package controllers

import (
	"Four/models"
	"Four/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(utils.BindError.Code, gin.H{"error": utils.BindError.Msg})
		return
	}
	comment.UserId = c.MustGet("user_id").(int)
	comment.PostId, _ = strconv.Atoi(c.Param("postid"))
	comment.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	if err := cc.DB.Create(&comment).Error; err != nil {
		c.JSON(utils.CreateError.Code, gin.H{"error": utils.CreateError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

func (cc *CommentController) GetComments(c *gin.Context) {
	var comments []models.Comment
	postid, _ := strconv.Atoi(c.Param("postid"))
	if err := cc.DB.Model(&models.Comment{}).Where("post_id=?", postid).Find(&comments).Error; err != nil {
		c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, comments)
	}
}
