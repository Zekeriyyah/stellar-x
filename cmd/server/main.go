package main

import (
	"os"

	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/routes"

	"github.com/Zekeriyyah/stellar-x/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	
	if os.Getenv("RENDER") == "true" {
		database.InitDB()

	} else {
		// separate database migration from binary-build in right order when using docker
		
		args := os.Args
		if len(args) > 1 && args[1] == "migrate" {
			
			database.InitDB()
			return
		}
	}
	
	r := gin.Default()

	// Setup all other routes 
	r = routes.SetupRouter(r) 

	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000" 
	}

	pkg.Info("Server starting on port " + port)

	if err := r.Run(":" + port); err != nil {
		pkg.Error(err, "failed to start server")
		return
	}
}
