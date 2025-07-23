package usecases

import (
	"errors"

	"LoginMechanism/internal/application/domain/entities" // User entity'sini kullanacağız
)

// Ortak hata tanımları. Bu hatalar use case'ler tarafından döndürülür
// ve adaptör katmanında (controller, repository) uygun HTTP/SQL hatalarına çevrilir.
var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("invalid email or password")
	ErrUserAlreadyExists     = errors.New("user with this email already exists")
	ErrPasswordHashFailed    = errors.New("failed to hash password")
	ErrTokenGenerationFailed = errors.New("failed to generate token")
)

// UserRepository, kullanıcı veritabanı işlemlerini soyutlayan arayüzdür.
// Application katmanı, veritabanının nasıl depolandığını bilmez; sadece bu arayüzü kullanır.
type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}

// PasswordHasher, şifre hashleme ve doğrulama işlemlerini soyutlayan arayüzdür.
// Application katmanı, hangi hash algoritmasının kullanıldığını (örn. bcrypt) bilmez.
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedPassword string) error
}

// TokenProvider, JWT token oluşturma ve doğrulama işlemlerini soyutlayan arayüzdür.
// Application katmanı, hangi token kütüphanesinin (örn. dgrijalva/jwt-go) kullanıldığını bilmez.
type TokenProvider interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (string, error) // Login use case'inde doğrudan kullanılmaz ama eklenecek middleware için faydalı
}
