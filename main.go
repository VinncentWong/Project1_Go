package main

import (
	"module/controller"
	"module/database"
	"module/envconfig"
)

func main() {
	envconfig.InitProperties()
	database.InitDb()
	controller.InitRoutes()
}
