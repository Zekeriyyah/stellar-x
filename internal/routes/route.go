// internal/routes/router.go
package routes

import (
	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/handlers"
	"github.com/Zekeriyyah/stellar-x/internal/middleware"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
		
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)
	walletRepo := repositories.NewWalletRepository(database.DB)
	balanceRepo := repositories.NewBalanceRepository(database.DB)
	transactionRepo := repositories.NewTransactionRepository(database.DB)
	auditLogRepo := repositories.NewAuditLogRepository(database.DB)

	// Initialize services

	userService := services.NewUserService(userRepo)
	walletService := services.NewWalletService(walletRepo, balanceRepo, userRepo)
	balanceService := services.NewBalanceService(balanceRepo)
	depositService := services.NewDepositService(balanceRepo, transactionRepo)
	auditLogService := services.NewAuditLogService(auditLogRepo)
	fxService := services.NewFXService()
	swapService := services.NewSwapService(balanceService, transactionRepo, fxService)
	transferService := services.NewTransferService(walletRepo, balanceService, transactionRepo, fxService)
	transactionService := services.NewTransactionService(transactionRepo, walletRepo)

	aiService := services.NewAIService(fxService)


	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	walletHandler := handlers.NewWalletHandler(walletService, userService)
	depositHandler := handlers.NewDepositHandler(depositService, walletService)
	auditLogHandler := handlers.NewAuditLogHandler(auditLogService)
	swapHandler := handlers.NewSwapHandler(swapService)
	transferHandler := handlers.NewTransferHandler(transferService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	aiHandler := handlers.NewAIHandler(aiService)

	
	// API v1


	r.POST("/api/signup", userHandler.Signup)
	r.GET("/api/login", userHandler.Login)

	api := r.Group("/api/v1")

	api.Use(middleware.AuthMiddleware())
	api.Use(middleware.AuditLogger(auditLogService)) // Log IP, device, browser, country

	{
		// User routes
		api.GET("/users/:userId", userHandler.GetByID)
		// api.GET("/users/email/:email", userHandler.GetByEmail)
		
		// Wallet routes
		api.POST("/wallet", walletHandler.CreateWallet)
		api.GET("/wallet/:userId", walletHandler.GetWallet)

		// Deposit routes
		api.POST("/deposit", depositHandler.Handle)

		// swap routes
		api.POST("/swap", swapHandler.Handle)

		// transfer routes
		api.POST("/transfer", transferHandler.Handle)

		// transaction routes
		api.GET("/transaction/:userId", transactionHandler.GetHistory)
		api.GET("/transactions", transactionHandler.GetRecent)

		// Audit log routes
		api.GET("/audit/:userId", auditLogHandler.GetAuditLogByUserID)

		// LLM chat bot 
		api.GET("/ask", aiHandler.Ask)
	}

	return r
}