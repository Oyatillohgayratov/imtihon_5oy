package handler

import (
	"api-gateway/api/models"
	checkemail "api-gateway/pkg/checkEmail"
	userpb "api-gateway/protos/user"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} userpb.RegisterUserResponse
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to register user"}
// @Router /api/users [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	r := checkemail.ValidateGmail(req.Email)
	if !r {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	resp, err := h.service.User().Register(context.Background(), &userpb.RegisterUserRequest{
		Username: req.Username, Password: req.Password, Email: req.Email,
	})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticate a user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} userpb.LoginUserResponse
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to login"}
// @Router /api/users/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.service.User().Login(context.Background(), &userpb.LoginUserRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
