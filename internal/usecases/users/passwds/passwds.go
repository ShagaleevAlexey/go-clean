package passwds

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordMatch = errors.New("hashedPassword is not the hash of the given password")

func GetPasswordHash(pass string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func MatchPasswordHash(hash string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return ErrPasswordMatch
	}

	return nil
}
