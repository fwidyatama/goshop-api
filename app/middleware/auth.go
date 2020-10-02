package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goshop-api/util"
	"net/http"
	"os"
	"strings"
)

func Auth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("signing method not matching")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token != nil && err == nil {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, util.FailResponse(http.StatusUnauthorized, "Please provide token"))
		c.Abort()
	}

}
