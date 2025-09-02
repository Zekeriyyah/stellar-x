package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/pkg"
	"github.com/gin-gonic/gin"
)

type UserHandlers interface {

	// User handlers
	Signup(c *gin.Context)
	Login(c *gin.Context)
	GetByID(c *gin.Context)
	GetByEmail(c *gin.Context)

	// Wallet Handlers
	CreateWallet(c *gin.Context)
	GetWallet(c *gin.Context)

	// Deposit Handler
	Handle(c *gin.Context)

	//Swap
	_Handle(c *gin.Context) 

	//Transfer Handler
	__Handle(c *gin.Context)

	//Transaction handler
	GetHistory(c *gin.Context) 
	GetRecent(c *gin.Context)

	//Audit Handler
	GetAuditLogByUserID(c *gin.Context) 
	
	//AI assistant Handler
	Ask(c *gin.Context)
}

func GetAllFxResp(c *gin.Context) {
	coinGecko, frankfurter := pkg.GetAllFxCurrencies()

	if coinGecko != nil || frankfurter!= nil {
		c.JSON(http.StatusOK, gin.H{"coingecko": coinGecko, "frankfurter": frankfurter})
		return
	}
	
	c.JSON(http.StatusGatewayTimeout, gin.H{"error": "failed to fetch apis"})
}