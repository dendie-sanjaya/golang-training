package handler

import (
	"net/http"
	"praisindo/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int
)

// Create User
func CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// Get User by ID
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// Update User
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == id {
			newUser := entity.User{
				ID:        id,
				Name:      user.Name,
				Email:     user.Email,
				Password:  u.Password,
				CreatedAt: u.CreatedAt,
				UpdatedAt: time.Now(),
			}

			users[i] = newUser
			c.JSON(http.StatusOK, newUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// Delete User
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// Get All Users
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
