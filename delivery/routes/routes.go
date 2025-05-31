package routes

import (
	"team-maker-api/delivery"
	"team-maker-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *delivery.Handlers) *gin.Engine {
	r := gin.Default()

	// Route ping untuk health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping team-maker-api v1.0",
		})
	})

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
