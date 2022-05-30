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

var mapData map[string]interface{} = make(map[string]interface{})

type AdminHandler struct{}

func SetDb(_db *gorm.DB) {
	db = _db
}

func (receiver *AdminHandler) Register(c *gin.Context) {
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

func (receiver *AdminHandler) Login(c *gin.Context) {
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
	Bcrypt.Matches(tempAdmin.Password, admin.Password)
	var token middleware.IJwt = Jwt
	jwtToken := token.GenerateToken(admin.ID, admin.Name, admin.Email)
	mapData["data"] = jwtToken
	response := util.Response{
		Success: true,
		Message: "Admin Authenticated",
		Data:    mapData,
	}
	c.JSON(http.StatusOK, response)
}
