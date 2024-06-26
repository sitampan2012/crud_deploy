package auth

import (
	"net/http"
	"simple-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	USER     = "admin"
	PASSWORD = "Password123!"
	SECRET   = "secret"
)

func LoginHandler(c *gin.Context) {
	var user models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user invalid",
		})
	} else if user.Password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password invalid",
		})
		return
	} else {

		//token
		// claim := jwt.StandardClaims{
		// 	ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		// 	Issuer:    "test",
		// 	IssuedAt:  time.Now().Unix(),
		// }

		claim := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
			Issuer:    "test",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(SECRET))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Succes",
			"token":   token,
		})
	}
}
