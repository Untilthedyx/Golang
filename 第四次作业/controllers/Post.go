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

type PostController struct {
	DB *gorm.DB
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(utils.BindError.Code, gin.H{"error": utils.BindError.Msg})
		return
	}

	post.UserId = c.MustGet("user_id").(int)
	post.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	// var user models.User
	// if err := pc.DB.Model(&models.User{}).Where("userid=?", post.UserId).First(&user); err != nil {
	// 	c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
	// 	return
	// }
	// if err := pc.DB.Model(&models.User{}).Where("userid=?", post.UserId).Update("posts", append(user.Posts, post)); err != nil {
	// 	c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
	// 	return
	// }
	if err := pc.DB.Create(&post).Error; err != nil {
		c.JSON(utils.CreateError.Code, gin.H{"error": utils.CreateError.Msg})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (pc *PostController) GetPost(c *gin.Context) {

	var post models.Post
	postid, _ := strconv.Atoi(c.Param("postid"))
	if err := pc.DB.Model(&models.Post{}).Where("post_id=?", postid).First(&post).Error; err != nil {
		c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}
func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := pc.DB.Find(&posts).Error; err != nil {
		c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(utils.BindError.Code, gin.H{"error": utils.BindError.Msg})
		return
	}
	post.PostId, _ = strconv.Atoi(c.Param("postid"))
	post.UpdatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	if err := pc.DB.Model(&models.Post{}).Where("post_id", post.PostId).Updates(post).Error; err != nil {
		c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func (pc *PostController) DeletePost(c *gin.Context) {
	var post models.Post
	post.PostId, _ = strconv.Atoi(c.Param("postid"))
	if err := pc.DB.Model(&models.Post{}).Where("post_id=?", post.PostId).Delete(&post).Error; err != nil {
		c.JSON(utils.DBError.Code, gin.H{"error": utils.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}
