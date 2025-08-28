package main

import (
	"os"

	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/routes"

	"github.com/Zekeriyyah/stellar-x/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB
	database.InitDB()
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Setup all other routes 
	r = routes.SetupRouter(r) 

	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000" 
	}

	pkg.Info("Server starting on port " + port)

	if err := r.Run(":" + port); err != nil {
		pkg.Error("Server failed to start", err)
	}
}
