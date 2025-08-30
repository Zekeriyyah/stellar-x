// internal/handlers/wallet_handler.go
package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	walletService *services.WalletService
	userService *services.UserService
}

type WalletInput struct {
		Email 	string `json:"email" binding:"required"`
		Label	string `json:"label"`
	}
	
func NewWalletHandler(w *services.WalletService, u *services.UserService) *WalletHandler {
	return &WalletHandler{
		walletService: w,
		userService: u,
	}
}



func(w *WalletHandler) CreateWallet(c *gin.Context) {
	
	input := &WalletInput{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or label"})
		return
	}

	// create wallet for user
	wallet , err := w.walletService.CreateWalletWithBalances(input.Email, input.Label)
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v created successful", wallet.Label),
		"body": gin.H{"wallet": wallet},
	})
}

func(w *WalletHandler) GetWallet(c *gin.Context) {
	
	userId := c.Param("userId")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	wallet, err := w.walletService.GetWalletByUserID(uint(userIdUint))
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "wallet not found"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"body": wallet,
		})
}

