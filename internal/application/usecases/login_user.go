package usecases

import (
	"fmt"
)

// LoginUserInput, LoginUser use case'ine gelen giriş verisini tanımlar.
type LoginUserInput struct {
	Email    string
	Password string
}

// LoginUserOutput, LoginUser use case'inden dönen çıkış verisini tanımlar.
type LoginUserOutput struct {
	Token string
}

// LoginUserUseCase, kullanıcı girişi işlemini gerçekleştiren arayüzdür.
// Bu da bir "interactor" veya "use case" olarak bilinir.
type LoginUserUseCase interface {
	Execute(input LoginUserInput) (*LoginUserOutput, error)
}

// loginUserInteractor, LoginUserUseCase arayüzünün somut implementasyonudur.
// Bağımlılıkları olarak UserRepository, PasswordHasher ve TokenProvider arayüzlerini alır.
type loginUserInteractor struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
	tokenProvider  TokenProvider
}

// NewLoginUserUseCase yeni bir loginUserInteractor instance'ı oluşturur.
// Bağımlılıklar (UserRepository, PasswordHasher, TokenProvider) dışarıdan enjekte edilir.
func NewLoginUserUseCase(
	userRepository UserRepository,
	passwordHasher PasswordHasher,
	tokenProvider TokenProvider,
) LoginUserUseCase {
	return &loginUserInteractor{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenProvider:  tokenProvider,
	}
}

// Execute metodu, LoginUser use case'inin ana iş mantığını içerir.
func (l *loginUserInteractor) Execute(input LoginUserInput) (*LoginUserOutput, error) {
	// 1. E-posta ve şifre geçerliliğini kontrol et
	if input.Email == "" || input.Password == "" {
		return nil, fmt.Errorf("email and password cannot be empty")
	}

	// 2. Kullanıcıyı e-posta ile veritabanından getir
	user, err := l.userRepository.GetUserByEmail(input.Email)
	if err == ErrUserNotFound {
		// Kullanıcı bulunamazsa, güvenlik nedeniyle "Geçersiz e-posta veya şifre" döndür
		return nil, fmt.Errorf("invalid email or password")
	}
	if err != nil {
		// Diğer veritabanı hatalarını handle et
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	// 3. Sağlanan şifreyi, veritabanındaki hashlenmiş şifre ile karşılaştır
	err = l.passwordHasher.CheckPassword(input.Password, user.Password)
	if err != nil {
		// Şifre eşleşmezse, "Geçersiz e-posta veya şifre" döndür
		return nil, fmt.Errorf("invalid email or password")
	}

	// 4. Kullanıcı kimliği doğrulandıktan sonra JWT token oluştur
	token, err := l.tokenProvider.GenerateToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// 5. Başarılı çıkış modelini döndür
	return &LoginUserOutput{Token: token}, nil
}
