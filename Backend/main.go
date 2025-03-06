package main

import (
	"backend/handler"
	sv_cfg "backend/servercfg"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//handlers User
	router.POST("/user", handler.CreateUser)
	router.GET("/users", handler.GetAllUsers)
	router.GET("/user/:username", handler.GetUserByUserName)
	router.PUT("/user/:username", handler.UpdateUser)
	router.DELETE("/user/:username", handler.DeleteUser)

	//handlers Product
	router.GET("/products", handler.GetAllProducts)
	router.POST("/product", handler.CreateProduct)
	router.GET("/products/:type", handler.GetProductByType)
	router.GET("/product/:id", handler.GetProductByID)
	router.PUT("/product/:id", handler.UpdateProduct)
	router.DELETE("/product/:id", handler.DeleteProduct)

	//Cargar Configuracion de la Base de Datos
	sv_cfg.InitDatabaseConfig()

	//Iniciar Servidor
	router.Run(sv_cfg.ServerHost + ":" + sv_cfg.ServerPort)
}
