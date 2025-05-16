package products

import (
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/products", GetAllProducts)
		products.POST("/product", CreateProduct)
		products.GET("/products/:type", GetProductByType)
		products.GET("/product/:id", GetProductByID)
		products.PUT("/product/:id", UpdateProduct)
		products.DELETE("/product/:id", DeleteProduct)
	}
}
