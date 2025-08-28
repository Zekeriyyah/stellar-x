package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	TransferService *services.TransferService
}

func NewTransferHandler(transferService *services.TransferService) *TransferHandler {
	return &TransferHandler{TransferService: transferService}
}

// Handle processes POST /api/v1/transfer
func (h *TransferHandler) Handle(c *gin.Context) {
	var input struct {
		SenderWalletID   uint  `json:"sender_wallet_id" binding:"required"`
		ReceiverWalletID uint  `json:"receiver_wallet_id" binding:"required"`
		FromCurrency     string  `json:"from_currency" binding:"oneof=cNGN cXAF USDx EURx"`
		ToCurrency		 string	 `json:"to_currency" binding:"oneof=cNGN cXAF USDx EURx"`
		Amount           float64 `json:"amount" binding:"gt=0"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Execute transfer
	transaction, err := h.TransferService.Transfer(input.SenderWalletID, input.ReceiverWalletID, input.FromCurrency, input.ToCurrency, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transfer successful",
		"transaction": transaction,
	})
}