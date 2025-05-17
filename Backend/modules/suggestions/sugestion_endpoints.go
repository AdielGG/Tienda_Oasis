package suggestions

import (
	"github.com/gin-gonic/gin"
)

func InitSuggestionRoutes(router *gin.Engine) {
	suggestions := router.Group("/suggestions")
	{
		suggestions.POST("/suggestion", CreateSuggestion)
		suggestions.GET("/suggestions", GetAllSuggestions)
		suggestions.GET("/suggestion/:id", GetSuggestion)
		suggestions.PUT("/suggestion/:id", UpdateSuggestion)
		suggestions.DELETE("/suggestion/:id", DeleteSuggestion)
	}
}
