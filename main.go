package main

import (
	"module/controller"
	"module/database"
)

func main() {
	database.InitDb()
	controller.InitRoutes()
}
