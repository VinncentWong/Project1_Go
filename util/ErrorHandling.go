package util

func HandlingError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
