package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wingobank/auth-service/internal/handlers"
	"github.com/wingobank/auth-service/internal/infrastructure/db"
	"github.com/wingobank/auth-service/internal/repositories"
	"github.com/wingobank/auth-service/internal/routes"
	"github.com/wingobank/auth-service/internal/services"
)

func main() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found. Using environment variables.")
	}

	// Connect to the db
	database := db.ConnectDB()

	// Set Gin mode (debug/release) from env
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	// Initialize router and set trusted proxies
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Dependency injection: repo -> service -> handler
	userRepo := repositories.NewUserRepository(database)
	authSvc := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authSvc)

	// Register routes under /auth
	authGroup := router.Group("/auth")
	routes.RegisterAuthRoutes(authGroup, authHandler)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Auth Service running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
