package auth

import (
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", CreateUser)
		auth.GET("/user/:username", GetUserByUserName)
		auth.POST("/login", Login)
		auth.GET("/users", GetAllUsers)
	}
}
