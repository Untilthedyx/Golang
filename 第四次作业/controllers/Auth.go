package controllers

import (
	"Four/config"
	"Four/models"
	"Four/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB     *gorm.DB
	Config *config.Config
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(utils.BindError.Code, gin.H{"error": utils.BindError.Msg})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(utils.HashError.Code, gin.H{"error": utils.HashError.Msg})
	}
	user.Password = string(hashedPassword)
	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(utils.CreateError.Code, gin.H{"error": utils.CreateError.Msg})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(utils.BindError.Code, gin.H{"error": utils.BindError.Msg})
		return
	}
	var dbUser models.User
	if err := ac.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(utils.UsernameOrPasswordError.Code, gin.H{"error": utils.UsernameOrPasswordError.Msg})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(utils.UsernameOrPasswordError.Code, gin.H{"error": utils.UsernameOrPasswordError.Msg})
		return
	}
	ac.Config = config.LoadConfig()
	token, err := utils.GenerateToken(dbUser, ac.Config.JWTSecret)
	if err != nil {
		c.JSON(utils.GenerateTokenError.Code, gin.H{"error": utils.GenerateTokenError.Msg})
		return
	}
	c.Header("Authorization", "Bearer"+token)
}
