package usecases

import (
	"errors"

	"LoginMechanism/internal/application/domain/entities" // We will use the User entity
)

// Common error definitions. These errors are returned by use cases
// and are translated into appropriate HTTP/SQL errors in the adapter layer (controller, repository).
var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("invalid email or password")
	ErrUserAlreadyExists     = errors.New("user with this email already exists")
	ErrPasswordHashFailed    = errors.New("failed to hash password")
	ErrTokenGenerationFailed = errors.New("failed to generate token")
)

// UserRepository is an interface that abstracts user database operations.
// The Application layer does not know how the database is stored; it only uses this interface.
type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}

// PasswordHasher is an interface that abstracts password hashing and verification operations.
// The Application layer does not know which hashing algorithm is used (e.g., bcrypt).
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedPassword string) error
}

// TokenProvider is an interface that abstracts JWT token generation and validation operations.
// The Application layer does not know which token library is used (e.g., dgrijalva/jwt-go).
type TokenProvider interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (string, error) // Not directly used in the Login use case but useful for future middleware
}
