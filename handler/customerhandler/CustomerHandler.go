package customerhandler

import (
	"module/entities"
	"module/entities/dto"
	"module/middleware"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var mapData map[string]interface{}
var jwt middleware.IJwt = new(middleware.Jwt)
var bcrypt util.Bcrypt

func SetDb(_db *gorm.DB) {
	db = _db
}

type CustomerHandler struct{}

func (*CustomerHandler) Register(c *gin.Context) {
	var bodyRegister dto.RegisterDto
	err := c.ShouldBindJSON(&bodyRegister)
	util.HandlingError(err)
	bodyRegister.Password = bcrypt.Encode(bodyRegister.Password)
	result := db.Create(&bodyRegister)
	util.HandlingError(result.Error)
	mapData["data"] = bodyRegister
	response := util.Response{
		Success: true,
		Message: "Success register Admin data! ",
		Data:    mapData,
	}
	c.JSON(http.StatusCreated, response)
}

func (*CustomerHandler) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	util.HandlingError(err)
	var customer entities.Customer
	result := db.Where("email = ?", loginDto.Email).Take(&customer)
	util.HandlingError(result.Error)
	isValid := bcrypt.Matches(loginDto.Password, customer.Password)
	if isValid {
		token := jwt.GenerateToken(customer.ID, customer.Name, customer.Email)
		mapData["data"] = customer
		mapData["token"] = token
		response := util.Response{
			Success: true,
			Message: "User authenticated! ",
			Data:    mapData,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := util.Response{
			Success: false,
			Message: "Unauthenticated",
			Data:    mapData,
		}
		c.JSON(http.StatusUnauthorized, response)
	}
}

func (*CustomerHandler) GetCustomerById(c *gin.Context) {
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var customer entities.Customer
	result := db.Where("id = ?", id).Take(&customer)
	util.HandlingError(result.Error)
	mapData["data"] = customer
	response := util.Response{
		Success: true,
		Message: "Success find book data! ",
		Data:    mapData,
	}
	c.JSON(http.StatusOK, response)
}

func (*CustomerHandler) UpdateCustomerById(c *gin.Context) {
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var bodyCustomer entities.Customer
	err := c.ShouldBindJSON(&bodyCustomer)
	util.HandlingError(err)
	result := db.Model(&bodyCustomer).Where("id = ?", id).Updates(&bodyCustomer)
	util.HandlingError(result.Error)
	mapData["data"] = bodyCustomer
	response := util.Response{
		Success: true,
		Data:    mapData,
		Message: "Success update customer data! ",
	}
	c.JSON(http.StatusOK, response)
}

func (*CustomerHandler) DeleteCustomerById(c *gin.Context) {
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var model entities.Customer
	result := db.Where("id = ?", id).Delete(&model)
	util.HandlingError(result.Error)
	response := util.Response{
		Message: "Success delete Customer Data",
		Data:    nil,
		Success: true,
	}
	c.JSON(http.StatusOK, response)
}

func (*CustomerHandler) GetAllBook(c *gin.Context) {
	var books []entities.Customer
	result := db.Find(&books)
	util.HandlingError(result.Error)
	mapData["data"] = books
	response := util.Response{
		Message: "Success get Customers Data",
		Data:    nil,
		Success: true,
	}
	c.JSON(http.StatusOK, response)
}

func (*CustomerHandler) GetBookById(c *gin.Context) {
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var book entities.Buku
	result := db.Where("id = ?", id).Take(&book)
	util.HandlingError(result.Error)
	mapData["data"] = book
	response := util.Response{
		Success: true,
		Message: "Success find book data! ",
		Data:    mapData,
	}
	c.JSON(http.StatusOK, response)
}
