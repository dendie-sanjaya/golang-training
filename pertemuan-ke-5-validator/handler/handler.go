package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"praisindo/entity"

	"praisindo/service"

	"github.com/gin-gonic/gin"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser menghandle permintaan untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var msg string
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "Field validation for 'Name' failed on the 'required' tag") {
			msg = msg + " nama kosong ~~~"
			//c.JSON(http.StatusBadRequest, gin.H{"error": "nama kosong"})
			//return
		}

		if strings.Contains(err.Error(), "Field validation for 'Email' failed on the 'required' tag") {
			msg = msg + " email tidak sesuai format ~~~"
			//c.JSON(http.StatusBadRequest, gin.H{"error": "Email tidak sesuai format"})
			//return
		}

		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	createdUser := h.userService.CreateUser(&user)
	c.JSON(http.StatusCreated, createdUser)
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser menghandle permintaan untuk mengupdate informasi user
func (h *UserHandler) UpdateUser(c *gin.Context) {
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

	updatedUser, err := h.userService.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser menghandle permintaan untuk menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
