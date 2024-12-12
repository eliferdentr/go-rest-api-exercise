package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
)

func HashPasswordArgon2id(password string) (string, string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if (err != nil) {
		fmt.Println(err)
		return "", "", err
	}
	hash := argon2.IDKey([]byte(password), salt, time, memory, uint8(threads), keyLen)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
 
	return b64Hash, b64Salt, nil
}

func CompareHashArgon2id(password, b64Salt, hashedPassword string) bool {
	salt, err := base64.RawStdEncoding.DecodeString(b64Salt)
	if err != nil {
		return false
	}
	hash := argon2.IDKey([]byte(password), salt, time, memory, uint8(threads), keyLen)
 
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return b64Hash == hashedPassword
}