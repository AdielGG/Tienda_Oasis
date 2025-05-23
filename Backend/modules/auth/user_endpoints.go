package auth

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)

		auth.GET("/user/:username", middleware.RequireAuth, GetUserByUserName)
		auth.PUT("/edit", middleware.RequireAuth, EditUser)
		auth.GET("/logout", middleware.RequireAuth, Logout)

		auth.GET("/users", middleware.RequireAuth, middleware.RequireAdmin, GetAllUsers)
		auth.POST("/createuser", CreateUser)
		auth.POST("/updateUser", middleware.RequireAuth, middleware.RequireAdmin, UpdateUser)

	}
}
