package handlers

import (
	"net/http"
	"strconv"

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
		SenderWalletID   string  `json:"senderWalletId" binding:"required"`
		ReceiverWalletID string  `json:"receiverWalletId" binding:"required"`
		FromCurrency     string  `json:"fromCurrency" binding:"oneof=cNGN cXAF USDx EURx"`
		ToCurrency		 string	 `json:"toCurrency" binding:"oneof=cNGN cXAF USDx EURx"`
		Amount           float64 `json:"amount" binding:"gt=0"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Parse wallet IDs
	senderID, err := strconv.ParseUint(input.SenderWalletID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sender wallet id"})
		return
	}

	receiverID, err := strconv.ParseUint(input.ReceiverWalletID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receiver wallet id"})
		return
	}

	// Execute transfer
	transaction, err := h.TransferService.Transfer(
		uint(senderID),
		uint(receiverID),
		input.FromCurrency,
		input.ToCurrency,
		input.Amount,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transfer successful",
		"transaction": transaction,
	})
}