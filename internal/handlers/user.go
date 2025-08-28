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

// Create handles POST /api/v1/users
// Used internally to create users for wallets
func (h *UserHandler) Create(c *gin.Context) {

	var input struct {
		Email string `json:"email" binding:"omitempty,email"`
		Phone string `json:"phone" binding:"omitempty"`
	}

	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := &models.User{}
	user.Email = input.Email
	user.Phone = input.Phone

	status, msg := h.UserService.CreateUser(user)
	if status != http.StatusOK {
		c.JSON(status, gin.H{"erro": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"body": gin.H{
			"email":   user.Email,
			"phone":   user.Phone,
		},
	})
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