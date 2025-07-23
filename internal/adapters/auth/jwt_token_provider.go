package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// jwtClaims JWT token içinde saklanacak özel claim'leri tanımlar.
// E-posta adresini ve standart JWT claim'lerini içerir.
type jwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// JWTTokenProvider, usecases.TokenProvider arayüzünün bir implementasyonudur.
// Tokenları imzalamak ve doğrulamak için bir gizli anahtara (secret key) ihtiyaç duyar.
type JWTTokenProvider struct {
	secretKey []byte
}

// NewJWTTokenProvider yeni bir JWTTokenProvider instance'ı oluşturur.
// Gizli anahtar dışarıdan enjekte edilir.
func NewJWTTokenProvider(secretKey []byte) *JWTTokenProvider {
	return &JWTTokenProvider{
		secretKey: secretKey,
	}
}

// GenerateToken, verilen e-posta için yeni bir JWT token oluşturur.
// Bu metot, usecases.TokenProvider arayüzünün GenerateToken metodunu uygular.
func (p *JWTTokenProvider) GenerateToken(email string) (string, error) {
	// Token'ın geçerlilik süresi: şu an + 5 dakika
	expirationTime := time.Now().Add(5 * time.Minute)

	// Custom claim'leri ve standart claim'leri içeren bir Claims objesi oluştur
	claims := &jwtClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(expirationTime.Unix())), // Token bitiş zamanı
			IssuedAt:  jwt.NewTime(float64(time.Now().Unix())),     // Token oluşturulma zamanı
			// NotBefore:   jwt.NewTime(float64(time.Now().Unix())), // Opsiyonel: Token ne zamandan önce geçerli değil
			// Audience:    []string{"your-app-audience"},          // Opsiyonel: Hedef kitle
			// Issuer:      "your-auth-service",                   // Opsiyonel: Tokenı veren
			// Subject:     email,                                 // Opsiyonel: Konu (genellikle kullanıcı kimliği)
		},
	}

	// Token'ı HS256 algoritması ve gizli anahtar ile imzala
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(p.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken, verilen token'ı doğrular ve geçerliyse içindeki claim'leri döndürür.
// Bu metot, usecases.TokenProvider arayüzünün ValidateToken metodunu uygular.
// Projenin mevcut scope'unda doğrudan kullanılmasa da, Auth middleware'i için gereklidir.
func (p *JWTTokenProvider) ValidateToken(tokenString string) (string, error) {
	claims := &jwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// HS256 algoritmasının kullanıldığından emin olun
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return p.secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("token validation failed: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	return claims.Email, nil
}
