package products

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/products", GetAllProducts)
		products.GET("/products/:type", GetProductByType)
		products.GET("/product/:id", GetProductByID)

		products.PUT("/product/:id", middleware.RequireAuth, middleware.RequireAdmin, UpdateProduct)
		products.DELETE("/product/:id", middleware.RequireAuth, middleware.RequireAdmin, DeleteProduct)
		products.POST("/product", middleware.RequireAuth, middleware.RequireAdmin, CreateProduct)
	}
}
