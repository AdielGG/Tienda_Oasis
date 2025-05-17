package resources

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitResourceRoutes(router *gin.Engine, resource *embed.FS) {

	//Serve img files
	router.GET("/img/:filename", func(c *gin.Context) {
		file, _ := resource.ReadFile("resource/img/" + c.Param("filename"))
		c.Data(http.StatusOK, "image/jpg", file)
	})

	//Serve programas files
	router.GET("/programs/:filename", func(c *gin.Context) {
		file, _ := resource.ReadFile("resource/programs/" + c.Param("filename"))
		c.Data(http.StatusOK, "application/zip", file)
	})

	//serve ebooks files
	router.GET("/pdf/:filename", func(c *gin.Context) {
		file, _ := resource.ReadFile("resource/pdf/" + c.Param("filename"))
		c.Data(http.StatusOK, "application/pdf", file)
	})

	//serve cursos.zip files
	router.GET("/cursos/:filename", func(c *gin.Context) {
		file, _ := resource.ReadFile("resource/cursos/" + c.Param("filename"))
		c.Data(http.StatusOK, "application/zip", file)
	})
}
