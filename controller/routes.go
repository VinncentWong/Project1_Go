package controller

import (
	"module/handler/adminhandler"
	"module/handler/customerhandler"
	"module/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	var jwt middleware.IJwt = new(middleware.Jwt)
	var adminHandler adminhandler.IAdminHandler = new(adminhandler.AdminHandler)
	var customerHandler customerhandler.ICustomerHandler = new(customerhandler.CustomerHandler)

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	admin := r.Group("/admin")
	admin.POST("/register", adminHandler.Register)
	admin.POST("/login", adminHandler.Login)
	admin.GET("/get/:id", jwt.ValidateToken(), adminHandler.GetAdminById)
	admin.DELETE("/delete/:id", jwt.ValidateToken(), adminHandler.DeleteAdminById)
	admin.PATCH("/update/:id", jwt.ValidateToken(), adminHandler.UpdateAdmin)
	admin.POST("/addbook", jwt.ValidateToken(), adminHandler.AddBook)
	admin.GET("/getbook/:id", jwt.ValidateToken(), adminHandler.GetBook)
	admin.PATCH("/updatebook/:id", jwt.ValidateToken(), adminHandler.UpdateBook)
	admin.DELETE("/deletebook/:id", jwt.ValidateToken(), adminHandler.DeleteBook)

	customer := r.Group("customer")
	customer.POST("/register", customerHandler.Register)
	customer.POST("/login", customerHandler.Login)
	customer.GET("/get/:id", jwt.ValidateToken(), customerHandler.GetCustomerById)
	customer.GET("/getbook", jwt.ValidateToken(), customerHandler.GetAllBook)
	customer.GET("/getbook/:id", jwt.ValidateToken(), customerHandler.GetBookById)
	customer.PATCH("/update/:id", jwt.ValidateToken(), customerHandler.UpdateCustomerById)
	customer.DELETE("/delete/:id", jwt.ValidateToken(), customerHandler.DeleteCustomerById)
	r.Run(":5000")
}
