// 用golang寫出一個UserController，包含以下功能：
// 1. 新增使用者
// 2. 列出所有使用者
// 3. 刪除使用者
// 4. 更新使用者

package controllers

import (
	"gotest/core"
	"gotest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create User
// @Produce json
// @Param user body models.User true "User"

// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	if err := core.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary List Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func ListUsers(c *gin.Context) {
	var users []models.User
	if err := core.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Delete User
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := core.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Update User
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	if err := core.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Search Users
// @Produce json
// @Param query query string true "Search Query"
// @Success 200 {array} models.User
// @Router /users/search [get]
func SearchUsers(c *gin.Context) {
	query := c.Query("query")
	var users []models.User
	if err := core.DB.Where("username LIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Count Users
// @Produce json
// @Success 200 {object} map[string]int
// @Router /users/count [get]
func CountUsers(c *gin.Context) {
	var count int64
	if err := core.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to count users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// @Summary Get User by Email
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} models.User
// @Router /users/email [get]
func GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	var user models.User
	if err := core.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Get User by Username
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} models.User
// @Router /users/username [get]
func GetUserByUsername(c *gin.Context) {
	username := c.Query("username")
	var user models.User
	if err := core.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Get User by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/id/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := core.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
