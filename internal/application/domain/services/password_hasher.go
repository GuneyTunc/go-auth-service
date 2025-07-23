package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" // We'll use the bcrypt library
)

// PasswordHasher is an interface that abstracts password hashing and verification operations.
// This interface will be used in the use case layer and represents a domain-specific service.
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedPassword string) error
}

// bcryptPasswordHasher is a concrete implementation of the PasswordHasher interface using the bcrypt library.
type bcryptPasswordHasher struct{}

// NewBcryptPasswordHasher creates and returns a new instance of bcryptPasswordHasher.
// This constructor function will be called within `main.go`.
func NewBcryptPasswordHasher() PasswordHasher {
	return &bcryptPasswordHasher{}
}

// HashPassword hashes the given plaintext password using bcrypt.
func (b *bcryptPasswordHasher) HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost uses a secure default cost value.
	// This value can be increased based on security requirements.
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPassword compares the given plaintext password with the hashed password.
// It returns nil if the passwords match, otherwise an error.
func (b *bcryptPasswordHasher) CheckPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// When bcrypt.ErrMismatchedHashAndPassword is returned, it means the passwords do not match.
		// Instead of using this error directly, we can return a more general error or nil.
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return fmt.Errorf("password does not match")
		}
		return fmt.Errorf("error comparing password hash: %w", err)
	}
	return nil // Passwords match
}
