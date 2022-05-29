package middleware

import (
	"module/util"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id uint, name string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})
	fixToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	util.HandlingError(err)
	return fixToken, err
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = token[7:]
		fixedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		util.HandlingError(err)
		_, ok := fixedToken.Claims.(jwt.MapClaims)
		if ok && fixedToken.Valid {
			c.Next()
			return
		} else {
			response := util.Response{
				Success: false,
				Message: "Unauthorizied!",
			}
			c.JSON(http.StatusForbidden, response)
		}
	}
}
