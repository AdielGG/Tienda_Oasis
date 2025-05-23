package suggestions

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitSuggestionRoutes(router *gin.Engine) {
	suggestions := router.Group("/suggestions")
	{
		suggestions.POST("/suggestion", CreateSuggestion)
		suggestions.GET("/suggestions", GetAllSuggestions)

		suggestions.GET("/suggestion/:id", middleware.RequireAuth, middleware.RequireAdmin, GetSuggestion)
		suggestions.PUT("/suggestion/:id", middleware.RequireAuth, middleware.RequireAdmin, UpdateSuggestion)
		suggestions.DELETE("/suggestion/:id", middleware.RequireAuth, middleware.RequireAdmin, DeleteSuggestion)
	}
}
