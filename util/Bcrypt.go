package util

import "golang.org/x/crypto/bcrypt"

type IBcrypt interface {
	Encode(password string) string
	Matches(rawPassword string, encodedPassword string) bool
}

type Bcrypt struct{}

func (receiver *Bcrypt) Encode(password string) string {
	passwordInByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordInByte, bcrypt.MinCost)
	HandlingError(err)
	return string(hashedPassword)
}

func (receiver *Bcrypt) Matches(rawPassword string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	isValid := err == nil
	if !isValid {
		HandlingError(err)
	}
	return isValid
}
