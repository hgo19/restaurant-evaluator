package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type EncrypterBcrypt struct{}

func (e *EncrypterBcrypt) HashPassword(password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("SALT_ROUNDS"))
	if err != nil {
		println("Error to convert .env salt_rounds to int")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltRound)
	return string(bytes), err
}
