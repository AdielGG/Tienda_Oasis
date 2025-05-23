package resources

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
	"time"

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

	/**********************************




	***********************************/

	upload := router.Group("/upload")
	//upload.Use(middleware.RequireAuth)
	//upload.Use(middleware.RequireAdmin)

	//upload images
	upload.POST("/img", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileExtension := file.Filename[strings.LastIndex(file.Filename, "."):]
		fileName := file.Filename[0:strings.LastIndex(file.Filename, ".")]
		file.Filename = fileName + fmt.Sprintf("%d", time.Now().Unix()) + fileExtension

		err = c.SaveUploadedFile(file, "./resource/img/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"image": file.Filename})
	})

	//Upload programas files
	upload.POST("/programs", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileExtension := file.Filename[strings.LastIndex(file.Filename, "."):]
		fileName := file.Filename[0:strings.LastIndex(file.Filename, ".")]
		file.Filename = fileName + fmt.Sprintf("%d", time.Now().Unix()) + fileExtension

		err = c.SaveUploadedFile(file, "./resource/programs/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"program": file.Filename})
	})

	//Upload ebooks files
	upload.POST("/ebooks", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileExtension := file.Filename[strings.LastIndex(file.Filename, "."):]
		fileName := file.Filename[0:strings.LastIndex(file.Filename, ".")]
		file.Filename = fileName + fmt.Sprintf("%d", time.Now().Unix()) + fileExtension

		err = c.SaveUploadedFile(file, "./resource/ebooks/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ebooks": file.Filename})
	})

	//Upload cursos files
	upload.POST("/cursos", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileExtension := file.Filename[strings.LastIndex(file.Filename, "."):]
		fileName := file.Filename[0:strings.LastIndex(file.Filename, ".")]
		file.Filename = fileName + fmt.Sprintf("%d", time.Now().Unix()) + fileExtension

		err = c.SaveUploadedFile(file, "./resource/cursos/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(file.Filename)
		c.JSON(http.StatusOK, gin.H{"curso": file.Filename})
	})
}
