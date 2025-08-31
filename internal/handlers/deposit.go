package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/Zekeriyyah/stellar-x/pkg"
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
	WalletId uint		`json:"wallet_id" binding:"required"`
}

func (d *DepositHandler) Handle(c *gin.Context) {
	input := &DepositInput{}
	
	if err := c.ShouldBindJSON(input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invaliid input"})
		return
	}

	// Get wallet by user id
	wallets, err := d.WalletService.GetWalletByUserID(input.UserId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "wallet with the user id not found"})
		return
	}

	requiredWallet := models.Wallet{}
	for i := 0; i < len(wallets); i++ {
		if input.WalletId == wallets[i].ID {
			requiredWallet = wallets[i]
			return
		}
	}

	if !pkg.IsNotEmpty(requiredWallet) {
		c.JSON(http.StatusNotFound, "user has no wallet of this id")
		return
	}

	// call deposit service 
	transaction, err := d.DepositService.Deposit(requiredWallet.ID, input.Currency, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deposit successful","transaction-details": transaction})
}