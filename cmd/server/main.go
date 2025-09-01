package main

import (
	"os"
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/routes"
	"github.com/gin-contrib/cors"

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

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://stellar-x-ui.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))


	// Landing page
	r.GET("/", func(c *gin.Context) {
    c.File("index.html")
	})
	r.StaticFile("/styles.css", "styles.css")

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
