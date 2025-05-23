package middleware

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
)

func RequireAuth(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	tokenString = tokenString[7:]
	fmt.Println(tokenString)
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		database.DB.First(&user, "id = ?", int(claims["sub"].(float64)))
		fmt.Println(time.Now().Unix(), "    ", int64(claims["exp"].(float64)))
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			c.AbortWithStatus(http.StatusUnauthorized)
		}
		fmt.Println(user)
		if user.Token == "" || user.Token != tokenString {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
