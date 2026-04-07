package auth

import "github.com/alexedwards/argon2id"

func HashPassword(plaintextPassword string) (string, error) {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func CheckPassword(plaintextPassword, hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plaintextPassword, hashedPassword)
	if err != nil {
		return false, err
	}

	return match, err
}
