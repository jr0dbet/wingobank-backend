package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wingobank/auth-service/config"
	"github.com/wingobank/auth-service/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found. Using environment variables.")
	}

	config.ConnectDB()

	env := os.Getenv("GIN_MODE")
	if env == "" {
		env = "debug"
	}
	gin.SetMode(env)

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.AuthRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Auth Service running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
