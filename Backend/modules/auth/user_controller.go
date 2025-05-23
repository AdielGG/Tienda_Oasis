package auth

import (
	"backend/config"
	"backend/database"
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hash)

	user.Role = "cliente"

	err = database.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Signup successfully", "user": user})
}

func Login(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := c.BindJSON(&loginRequest)

	fmt.Println(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err = database.DB.Where("username = ?", loginRequest.Username).First(&user).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario o contraseña incorrectos"})
		return

	}

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al Crear el token"})
		return
	}
	user.Token = tokenString
	database.DB.Save(&user)

	user.Password = ""

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", "", -1, "/", "", false, false)

	user, _ := c.Get("user")
	c.Set("user", nil)
	var usuario models.User
	usuario = user.(models.User)

	usuario.Token = ""
	database.DB.Save(&usuario)
	c.JSON(http.StatusOK, gin.H{"OK": "Logout successfully"})

}

func EditUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userDB models.User
	database.DB.First(&userDB, "id = ?", user.ID)

	if userDB.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario no encontrado"})
		return
	}

	user.Role = userDB.Role
	user.Email = userDB.Email
	user.Username = userDB.Username

	err = database.DB.Save(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Edit successfully"})
}

func GetAllUsers(c *gin.Context) {

	var users []models.User

	err := database.DB.Find(&users).Error

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, users)
}

func GetUserByUserName(c *gin.Context) {
	username := c.Param("username")

	var user models.User

	err := database.DB.First(&user, username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, user)

}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hash)

	err = database.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Signup successfully"})
}

func UpdateUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userDB models.User
	database.DB.First(&userDB, "id = ?", user.ID)

	if userDB.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario no encontrado"})
		return
	}

	err = database.DB.Save(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Edit successfully"})
}
