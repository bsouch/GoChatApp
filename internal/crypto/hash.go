package crypto

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) ([]byte, error) {
	hashFunction := sha256.New()
	_, err := hashFunction.Write([]byte(password))
	if err != nil {
		return nil, err
	}

	fmt.Printf("%x", hashFunction)
	return hashFunction.Sum(nil), nil
}

func DoPasswordsMatch(password string, storedPasswordHash []byte) (bool, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hashedPassword, storedPasswordHash), nil
}
