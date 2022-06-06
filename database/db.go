package database

import (
	"fmt"
	"module/entities"
	"module/handler/adminhandler"
	"module/handler/customerhandler"
	"module/util"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	fmt.Println(os.Getenv("DNS"))
	_db, err := gorm.Open(mysql.Open(os.Getenv("DNS")), &gorm.Config{})
	util.HandlingError(err)
	db = _db
	err = db.AutoMigrate(&entities.Admin{}, &entities.Buku{}, &entities.Customer{})
	util.HandlingError(err)
	adminhandler.SetDb(db)
	customerhandler.SetDb(db)
	return db
}
