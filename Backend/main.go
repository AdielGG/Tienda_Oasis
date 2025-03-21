package main

import (
	conf "backend/config"
	"backend/database"
	"backend/handler"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	//handlers User
	router.POST("/register", handler.CreateUser)
	router.GET("/user/:username", handler.GetUserByUserName)
	router.POST("/login", handler.Login)
	// router.GET("/users", handler.GetAllUsers)
	// router.PUT("/user/:username", handler.UpdateUser)
	// router.DELETE("/user/:username", handler.DeleteUser)

	//handlers Product
	router.GET("/products", handler.GetAllProducts)
	router.POST("/product", handler.CreateProduct)
	router.GET("/products/:type", handler.GetProductByType)
	router.GET("/product/:id", handler.GetProductByID)
	router.PUT("/product/:id", handler.UpdateProduct)
	router.DELETE("/product/:id", handler.DeleteProduct)

	//handlers Sugestion

	router.POST("/suggestion", handler.CreateSugestion)
	router.GET("/suggestions", handler.GetAllSugestions)
	router.GET("/suggestion/:id", handler.GetSugestion)
	router.PUT("/suggestion/:id", handler.UpdateSugestion)
	router.DELETE("/suggestion/:id", handler.DeleteSugestion)

	//Cargar Configuracion de la Base de Datos
	// namedb, err := ioutil.ReadFile("basedatos.txt")
	// password, err := ioutil.ReadFile("password.txt")
	// if err != nil {
	// 	panic(err)
	// }

	conf.InitDatabaseConfig()
	// conf.InitDatabaseConfig(string(namedb), string(password))
	database.CreateDB()

	//Iniciar Servidor
	router.Run(conf.ServerHost + ":" + conf.ServerPort)
	fmt.Scanln()
}
