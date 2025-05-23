package middleware

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin(c *gin.Context) {

	admin, _ := c.Get("user")
	if admin.(models.User).Role != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
