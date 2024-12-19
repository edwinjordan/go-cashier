package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)

	return string(pwd)
}
