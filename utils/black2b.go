package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

func GenerateTextToHash(password string) string {
	passwordBytes := []byte(password)

	// Create a Blake2b hash with a 32-byte output size
	hash, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}

	// Write the input text bytes to the hash
	hash.Write(passwordBytes)

	hash32 := hash.Sum(nil)

	return hex.EncodeToString(hash32)
}
