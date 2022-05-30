package middleware

import "github.com/gin-gonic/gin"

type IJwt interface {
	GenerateToken(id uint, name string, email string) string
	ValidateToken() gin.HandlerFunc
}
