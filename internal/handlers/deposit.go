package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type DepositHandler struct {
	DepositService *services.DepositService
}

func NewDepositHandler(d *services.DepositService) *DepositHandler {
	return &DepositHandler{
		DepositService: d,
	}
}

type DepositInput struct{
	WalletID uint		`json:"wallet_id" binding:"required"`
	Currency string 	`json:"currency" binding:"required"`
	Amount float64		`json:"amount" binding:"required"`
}

func (d *DepositHandler) Handle(c *gin.Context) {
	input := &DepositInput{}
	
	if err := c.ShouldBindJSON(input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invaliid input"})
		return
	}

	
	// call deposit service 
	transaction, err := d.DepositService.Deposit(input.WalletID, input.Currency, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit",
		"body": gin.H{"transaction-details": transaction},
	})
}