package handlers

import (
	"log"
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type DepositHandler struct {
	DepositService *services.DepositService
	WalletService *services.WalletService
}

func NewDepositHandler(d *services.DepositService, w *services.WalletService) *DepositHandler {
	return &DepositHandler{
		DepositService: d,
		WalletService: w,
	}
}

type DepositInput struct{
	UserId uint		`json:"user_id" binding:"required"`
	Currency string 	`json:"currency" binding:"required"`
	Amount float64		`json:"amount" binding:"required"`
}

func (d *DepositHandler) Handle(c *gin.Context) {
	input := &DepositInput{}
	
	if err := c.ShouldBindJSON(input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invaliid input"})
		return
	}

	// Get wallet by user id
	wallet, err := d.WalletService.GetWalletByUserID(input.UserId)
	log.Print(wallet)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "wallet with the user id not found"})
		return
	}

	// call deposit service 
	transaction, err := d.DepositService.Deposit(wallet.ID, input.Currency, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit successful",
		"body": gin.H{"transaction-details": transaction},
	})
}