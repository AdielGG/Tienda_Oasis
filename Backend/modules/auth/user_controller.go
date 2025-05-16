package auth

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	user.ID = 43

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	role := user.Role
	if len(role) < 1 {
		user.Role = "cliente"
	}

	err = database.CreateUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Logued in successfully"})
}

func GetUserByUserName(c *gin.Context) {
	user, err := database.GetUserByUserName(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username := loginRequest.Username
	password := loginRequest.Password

	user, err := database.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": user.Username,
	// 	"id":       user.ID,
	// })
	// tokenString, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
