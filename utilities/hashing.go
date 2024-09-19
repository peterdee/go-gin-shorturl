package utilities

import "github.com/alexedwards/argon2id"

func CompareHashWithPlaintext(hashed, plaintext string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plaintext, hashed)
}

func CreateHash(plaintext string) (string, error) {
	return argon2id.CreateHash(plaintext, argon2id.DefaultParams)
}
