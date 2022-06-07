package database

import (
	"fmt"
	"module/entities"
	"module/handler/adminhandler"
	"module/handler/customerhandler"
	"module/util"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

func InitDbWithSupabase() *gorm.DB {
	fmt.Println("Get Supabase Data")
	dsn := fmt.Sprintf(
		"user=%s "+
			"password=%s "+
			"host=%s "+
			"TimeZone=Asia/Singapore "+
			"port=%s "+
			"dbname=%s",
		os.Getenv("SUPABASE_USER"),
		os.Getenv("SUPABASE_PASSWORD"),
		os.Getenv("SUPABASE_HOST"),
		os.Getenv("SUPABASE_PORT"),
		os.Getenv("SUPABASE_DB_NAME"))
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.HandlingError(err)
	err = _db.AutoMigrate(&entities.Admin{}, &entities.Buku{}, &entities.Customer{})
	util.HandlingError(err)
	db = _db
	adminhandler.SetDb(db)
	customerhandler.SetDb(db)
	return db
}
