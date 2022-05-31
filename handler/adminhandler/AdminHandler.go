package adminhandler

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
var Bcrypt *util.Bcrypt = new(util.Bcrypt)
var Jwt *middleware.Jwt = new(middleware.Jwt)

type AdminHandler struct{}

func SetDb(_db *gorm.DB) {
	db = _db
}

func (*AdminHandler) Register(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	var registerDto dto.RegisterDto
	err := c.ShouldBindJSON(&registerDto)
	util.HandlingError(err)
	admin := entities.Admin{
		Email:    registerDto.Email,
		Password: registerDto.Password,
		Name:     registerDto.Name,
	}
	admin.Password = Bcrypt.Encode(admin.Password)
	result := db.Create(&admin)
	util.HandlingError(result.Error)
	mapData["data"] = admin
	response := util.Response{
		Success: true,
		Message: "Success register Admin data! ",
		Data:    mapData,
	}
	c.JSON(http.StatusCreated, response)
}

func (*AdminHandler) Login(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	var loginDto dto.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	util.HandlingError(err)
	tempAdmin := entities.Admin{
		Email:    loginDto.Email,
		Password: loginDto.Password,
	}
	var admin entities.Admin
	result := db.Where("email = ?", tempAdmin.Email).Take(&admin)
	util.HandlingError(result.Error)
	isValid := Bcrypt.Matches(tempAdmin.Password, admin.Password)
	if isValid {
		var token middleware.IJwt = Jwt
		jwtToken := token.GenerateToken(admin.ID, admin.Name, admin.Email)
		mapData["data"] = admin
		mapData["token"] = jwtToken
		response := util.Response{
			Success: true,
			Message: "Admin Authenticated",
			Data:    mapData,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := util.Response{
			Success: false,
			Message: "Unauthorizied",
			Data:    nil,
		}
		c.JSON(http.StatusUnauthorized, response)
	}
}

func (*AdminHandler) GetAdminById(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var admin entities.Admin
	result := db.Where("id = ?", id).Take(&admin)
	util.HandlingError(result.Error)
	mapData["data"] = admin
	response := util.Response{
		Success: true,
		Message: "Admin data found! ",
		Data:    mapData,
	}
	c.JSON(http.StatusOK, response)
}

func (*AdminHandler) DeleteAdminById(c *gin.Context) {
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var admin entities.Admin
	result := db.Where("id = ?", id).Delete(&admin)
	util.HandlingError(result.Error)
	response := util.Response{
		Message: "Success delete Admin Data",
		Data:    nil,
		Success: true,
	}
	c.JSON(http.StatusOK, response)
}

func (*AdminHandler) UpdateAdmin(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	id, idExist := c.Params.Get("id")
	if !idExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var bodyAdmin entities.Admin
	err := c.ShouldBindJSON(&bodyAdmin)
	util.HandlingError(err)
	if len(bodyAdmin.Password) > 0 {
		bodyAdmin.Password = Bcrypt.Encode(bodyAdmin.Password)
	}
	admin := entities.Admin{
		Name:     bodyAdmin.Name,
		Password: bodyAdmin.Password,
		Email:    bodyAdmin.Email,
	}
	result := db.Model(&admin).Where("id = ?", id).Updates(&admin)
	util.HandlingError(result.Error)
	mapData["data"] = bodyAdmin
	response := util.Response{
		Success: true,
		Data:    mapData,
		Message: "Success update Admin data! ",
	}
	c.JSON(http.StatusOK, response)
}

func (*AdminHandler) AddBook(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	var book entities.Buku
	err := c.ShouldBindJSON(&book)
	util.HandlingError(err)
	result := db.Create(&book)
	util.HandlingError(result.Error)
	mapData["data"] = book
	response := util.Response{
		Message: "Success add book data! ",
		Success: true,
		Data:    mapData,
	}
	c.JSON(http.StatusCreated, response)
}

func (*AdminHandler) GetBook(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
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

func (*AdminHandler) DeleteBook(c *gin.Context) {
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
	result := db.Where("id = ?", id).Delete(&book)
	util.HandlingError(result.Error)
	response := util.Response{
		Message: "Success delete Book Data",
		Data:    nil,
		Success: true,
	}
	c.JSON(http.StatusOK, response)
}

func (*AdminHandler) UpdateBook(c *gin.Context) {
	var mapData map[string]interface{} = make(map[string]interface{})
	id, isIdExist := c.Params.Get("id")
	if !isIdExist {
		response := util.Response{
			Success: false,
			Message: "Error when system was querying in database",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
	}
	var book entities.Buku
	err := c.ShouldBindJSON(&book)
	util.HandlingError(err)
	result := db.Model(&book).Where("id = ?", id).Updates(&book)
	util.HandlingError(result.Error)
	mapData["data"] = book
	response := util.Response{
		Success: true,
		Data:    mapData,
		Message: "Success update Book data! ",
	}
	c.JSON(http.StatusOK, response)
}
