package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	// "gorm.io/gorm"
	"gotest/core"
	"gotest/models"
	"net/http"
	"time"
)

var jwtSecret = []byte("secret")

// @Summary Login
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	var dbUser models.User
	if err := core.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	claims := jwt.MapClaims{
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtSecret)

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

// @Summary Protected
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /protected [get]
func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized"})
}
