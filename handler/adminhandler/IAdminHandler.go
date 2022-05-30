package adminhandler

import "github.com/gin-gonic/gin"

type IAdminHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetAdminById(c *gin.Context)
	DeleteAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	AddBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
	GetBook(c *gin.Context)
}
