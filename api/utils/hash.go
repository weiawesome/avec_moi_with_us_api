package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GenerateHashPassword(password string, salt []byte) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(password))
	sha256Hash.Write(salt)
	hashedPassword := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}

func ValidatePassword(password string, targetPassword string, salt []byte) bool {
	return GenerateHashPassword(password, salt) == targetPassword
}

func GeneralHash(val string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(val))
	hashedValue := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}
