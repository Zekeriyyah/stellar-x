package handlers

import (
	"net/http"
	"strconv"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: service}
}

// GetHistory handles GET /api/v1/transactions/:userId
func (h *TransactionHandler) GetHistory(c *gin.Context) {
	userIDStr := c.Param("userId")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	transactions, err := h.TransactionService.GetByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no transactions found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":      userID,
		"transactions": transactions,
	})
}


// GetRecent handles GET /api/v1/transactions
// Query params: page, page_size
func (h *TransactionHandler) GetRecent(c *gin.Context) {
	// Parse page
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// Parse page_size
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Get transactions
	result, err := h.TransactionService.GetRecent(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, result)
	
}