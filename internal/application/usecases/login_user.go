package usecases

import (
	"fmt"
)

// LoginUserInput defines the input data for the LoginUser use case.
type LoginUserInput struct {
	Email    string
	Password string
}

// LoginUserOutput defines the output data returned by the LoginUser use case.
type LoginUserOutput struct {
	Token string
}

// LoginUserUseCase is the interface that performs the user login operation.
// This is also known as an "interactor" or "use case".
type LoginUserUseCase interface {
	Execute(input LoginUserInput) (*LoginUserOutput, error)
}

// loginUserInteractor is the concrete implementation of the LoginUserUseCase interface.
// It takes UserRepository, PasswordHasher, and TokenProvider interfaces as its dependencies.
type loginUserInteractor struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
	tokenProvider  TokenProvider
}

// NewLoginUserUseCase creates and returns a new instance of loginUserInteractor.
// The dependencies (UserRepository, PasswordHasher, TokenProvider) are injected from the outside.
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

// Execute method contains the main business logic for the LoginUser use case.
func (l *loginUserInteractor) Execute(input LoginUserInput) (*LoginUserOutput, error) {
	// 1. Validate email and password inputs.
	if input.Email == "" || input.Password == "" {
		return nil, fmt.Errorf("email and password cannot be empty")
	}

	// 2. Retrieve the user from the database by email.
	user, err := l.userRepository.GetUserByEmail(input.Email)
	if err == ErrUserNotFound {
		// If the user is not found, return a generic "invalid email or password" for security reasons.
		return nil, fmt.Errorf("invalid email or password")
	}
	if err != nil {
		// Handle other database retrieval errors.
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	// 3. Compare the provided password with the hashed password from the database.
	err = l.passwordHasher.CheckPassword(input.Password, user.Password)
	if err != nil {
		// If passwords do not match, return "invalid email or password".
		return nil, fmt.Errorf("invalid email or password")
	}

	// 4. Generate a JWT token after successful user authentication.
	token, err := l.tokenProvider.GenerateToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// 5. Return the successful output model.
	return &LoginUserOutput{Token: token}, nil
}
