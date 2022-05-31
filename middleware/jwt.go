package middleware

import (
	"module/util"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct{}

func (receiver *Jwt) GenerateToken(id uint, name string, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})
	fixToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	util.HandlingError(err)
	return fixToken
}

func (receiver *Jwt) ValidateToken() gin.HandlerFunc {
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

func (*Jwt) GenerateRefreshToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	fixedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	util.HandlingError(err)
	return fixedToken
}

func (*Jwt) ValidateRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = token[7:]
		fixedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		util.HandlingError(err)
		claims, ok := fixedToken.Claims.(jwt.MapClaims)
		id, ok := claims["id"].(uint)
		if !ok {
			response := util.Response{
				Success: false,
				Message: "Error when do assertion type!",
				Data:    nil,
			}
			c.JSON(http.StatusInternalServerError, response)
		}
		if ok && fixedToken.Valid {
			mapData := make(map[string]interface{})
			refreshToken := new(Jwt).GenerateRefreshToken(id)
			mapData["refreshToken"] = refreshToken
			mapData["jwtToken"] = new(Jwt).GenerateToken(id, "", "")
			c.Next()
		} else {
			response := util.Response{
				Success: false,
				Message: "Unauthorizied!",
			}
			c.JSON(http.StatusForbidden, response)
		}
	}
}
