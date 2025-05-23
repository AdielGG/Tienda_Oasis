package main

import (
	"backend/config"
	"backend/database"
	auth "backend/modules/auth"
	products "backend/modules/product"
	suggestions "backend/modules/suggestions"
	resources "backend/resource"

	"embed"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed resource
var resource embed.FS

func main() {

	//Iniciar router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	//Serve resource files
	router.StaticFS("/resource", http.FS(resource))
	resources.InitResourceRoutes(router, &resource)

	//Endpoints User
	auth.InitAuthRoutes(router)

	//Endpoints Product
	products.InitProductRoutes(router)

	//Endpoints Suggestion
	suggestions.InitSuggestionRoutes(router)

	//Iniciar Base de Datos
	database.InitDB()

	//Iniciar Servidor
	router.Run(config.ServerHost + ":" + config.ServerPort)
	fmt.Scanln()
}
