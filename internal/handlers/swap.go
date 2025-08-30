package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type SwapHandler struct {
	SwapService *services.SwapService
}

func NewSwapHandler(swapService *services.SwapService) *SwapHandler {
	return &SwapHandler{SwapService: swapService}
}

// Handle processes POST /api/v1/swap
func (h *SwapHandler) Handle(c *gin.Context) {
	var input struct {
		WalletID     uint    `json:"walletId" binding:"required"`
		FromCurrency string  `json:"fromCurrency" binding:"oneof=cNGN cXAF USDx EURx"`
		ToCurrency   string  `json:"toCurrency" binding:"oneof=cNGN cXAF USDx EURx"`
		Amount       float64 `json:"amount" binding:"gt=0"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Execute swap
	transaction, err := h.SwapService.Swap(input.WalletID, input.FromCurrency, input.ToCurrency, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Swap successful",
		"transaction": transaction,
	})
}