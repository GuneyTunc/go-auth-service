package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// jwtClaims defines the custom claims to be stored inside the JWT token.
// It includes the user's email address and standard JWT claims.
type jwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// JWTTokenProvider is an implementation of the usecases.TokenProvider interface.
// It requires a secret key to sign and validate tokens.
type JWTTokenProvider struct {
	secretKey []byte
}

// NewJWTTokenProvider creates and returns a new instance of JWTTokenProvider.
// The secret key is injected from the outside.
func NewJWTTokenProvider(secretKey []byte) *JWTTokenProvider {
	return &JWTTokenProvider{
		secretKey: secretKey,
	}
}

// GenerateToken creates a new JWT token for the given email.
// This method implements the GenerateToken method of the usecases.TokenProvider interface.
func (p *JWTTokenProvider) GenerateToken(email string) (string, error) {
	// Token expiration time: current time + 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create a Claims object containing custom claims and standard claims
	claims := &jwtClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(expirationTime.Unix())), // Token expiration time
			IssuedAt:  jwt.NewTime(float64(time.Now().Unix())),     // Token issuance time
			// NotBefore:   jwt.NewTime(float64(time.Now().Unix())), // Optional: Token is not valid before this time
			// Audience:    []string{"your-app-audience"},           // Optional: Intended audience
			// Issuer:      "your-auth-service",                     // Optional: Token issuer
			// Subject:     email,                                   // Optional: Subject (usually user ID)
		},
	}

	// Sign the token using the HS256 algorithm and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(p.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates the given token and returns its claims if valid.
// This method implements the ValidateToken method of the usecases.TokenProvider interface.
// Although not directly used within the current scope of the project, it's essential for Auth middleware.
func (p *JWTTokenProvider) ValidateToken(tokenString string) (string, error) {
	claims := &jwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the HS256 algorithm is used
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

	return claims.Email, nil // Return the email on successful validation
}
