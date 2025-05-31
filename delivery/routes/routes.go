package routes

import (
	"team-maker-api/delivery"
	"team-maker-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *delivery.Handlers) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Auth.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// user := protected.Group("/user")
		// {
		// 	user.GET("/profile", h.User.GetProfile)
		// 	user.PUT("/profile", h.User.UpdateProfile)
		// }
	}

	return r
}
