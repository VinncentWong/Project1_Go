package adminhandler

import "github.com/gin-gonic/gin"

type IAdminHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetAdminById(c *gin.Context)
	DeleteAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
}
