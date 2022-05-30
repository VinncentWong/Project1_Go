package customerhandler

import "github.com/gin-gonic/gin"

type ICustomerHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetCustomerById(c *gin.Context)
	GetAllBook(c *gin.Context)
	GetBookById(c *gin.Context)
	UpdateCustomerById(c *gin.Context)
	DeleteCustomerById(c *gin.Context)
}
