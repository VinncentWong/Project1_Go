package envconfig

import (
	"module/util"

	"github.com/joho/godotenv"
)

func InitProperties() {
	err := godotenv.Load()
	util.HandlingError(err)
}
