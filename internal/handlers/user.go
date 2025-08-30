package handlers

import (
	"net/http"
	"strconv"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Signup handles POST /api/v1/user
func(u *UserHandler) Signup(c *gin.Context) {

	user := &models.User{}

	// bind JSON request to user struct
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	status, msg := u.UserService.CreateUser(user)
	if status != http.StatusOK {
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user": gin.H{
			"id": user.ID,
			"email": user.Email,
			"created_at": user.CreatedAt,
		},
	})
}

// Login: handles user login into account
func(u *UserHandler) Login(c *gin.Context) {
	// get input from request
	// verify email
	// verify password
	// if successful, generate token
	// send response with the token

	var input struct {
		Email    string `json:"email"`
		Password string `json:"-"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tokenStr, status, msg := u.UserService.Login(input.Email, input.Password)
	if status != http.StatusOK {
		c.JSON(status, gin.H{"error": msg})
		return
	}	
	
	c.JSON(http.StatusOK, gin.H{"message": msg, "token": tokenStr})
}

// GetUserByID handles GET /api/v1/users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, status, msg := h.UserService.GetUserByID(uint(id))
	if status != http.StatusOK {
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetByEmail handles GET /api/v1/users/email/:email
func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user, status, msg := h.UserService.GetUserByEmail(email)
	if status != http.StatusOK {
		c.JSON(status, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}