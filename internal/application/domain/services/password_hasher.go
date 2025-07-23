package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" // bcrypt kütüphanesini kullanacağız
)

// PasswordHasher, şifre hashleme ve doğrulama işlemlerini soyutlayan arayüzdür.
// Bu arayüz, usecase katmanında kullanılacak ve domain'e özgü bir hizmeti temsil eder.
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedPassword string) error
}

// bcryptPasswordHasher, PasswordHasher arayüzünün bcrypt kütüphanesini kullanan somut implementasyonudur.
type bcryptPasswordHasher struct{}

// NewBcryptPasswordHasher yeni bir bcryptPasswordHasher instance'ı oluşturur.
// Bu constructor fonksiyonu, `main.go` içinde çağrılacaktır.
func NewBcryptPasswordHasher() PasswordHasher {
	return &bcryptPasswordHasher{}
}

// HashPassword, verilen düz metin şifreyi bcrypt kullanarak hashler.
func (b *bcryptPasswordHasher) HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost, güvenli bir varsayılan maliyet değerini kullanır.
	// Güvenlik gereksinimlerine göre bu değer artırılabilir.
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}
	return string(hashedBytes), nil
}

// CheckPassword, verilen düz metin şifreyi hashlenmiş şifreyle karşılaştırır.
// Şifreler eşleşirse nil, aksi takdirde bir hata döndürür.
func (b *bcryptPasswordHasher) CheckPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// bcrypt.ErrMismatchedHashAndPassword hatası döndüğünde, şifrelerin eşleşmediği anlamına gelir.
		// Bu hatayı doğrudan kullanmak yerine daha genel bir hata veya nil döndürebiliriz.
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return fmt.Errorf("password does not match")
		}
		return fmt.Errorf("error comparing password hash: %w", err)
	}
	return nil // Şifreler eşleşiyor
}
