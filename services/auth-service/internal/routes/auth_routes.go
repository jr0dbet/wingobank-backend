package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wingobank/auth-service/internal/handlers"
	"github.com/wingobank/auth-service/internal/middleware"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, h *handlers.AuthHandler) {
	rg.POST("/register", h.Register)
	rg.POST("/login", h.Login)

	// Protected routes with JWT
	protected := rg.Group("/")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/profile", h.Profile)
	}
}
