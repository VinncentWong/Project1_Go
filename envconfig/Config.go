package envconfig

import (
	"module/util"

	"github.com/joho/godotenv"
)

type Properties struct {
	SECRET_KEY string
}

func (receiver *Properties) InitProperties() {
	err := godotenv.Load()
	util.HandlingError(err)
}
