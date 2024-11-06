package passhash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PassHash contains methods for generating and validating password hashes.
type PassHash struct {
	cost int
}

// New returns a PasswordHasher instance.
func New() *PassHash {
	return &PassHash{
		cost: bcrypt.DefaultCost,
	}
}

// PlainTextToHash creates a password hash using bcrypt.
func (ph PassHash) PlainTextToHash(plaintextPassword string) (string, error) {
	hashByteArr, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), ph.cost)
	if err != nil {
		return "", fmt.Errorf("error generating password hash: %w", err)
	}

	return string(hashByteArr), nil
}

// ComparePassToHash checks if a plain text password is valid for a password hash (check if password is correct).
func (ph PassHash) ComparePassToHash(plaintextPassword, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintextPassword))
}
