package utils

import "golang.org/x/crypto/bcrypt"

func CreateHash(val string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	return string(hashed), err
}

func CompareHash(val string, hash string) error {

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(val))
}
