package usecases

import (
	"fmt"

	"LoginMechanism/internal/application/domain/entities" // We'll use the User entity
)

// RegisterUserInput defines the input data required for the RegisterUser use case.
type RegisterUserInput struct {
	Email    string
	Password string
}

// RegisterUserUseCase is the interface that defines the contract for registering a new user.
// In Clean Architecture, this is often referred to as an "interactor" or "use case".
type RegisterUserUseCase interface {
	Execute(input RegisterUserInput) error
}

// registerUserInteractor is the concrete implementation of the RegisterUserUseCase interface.
// It depends on the UserRepository and PasswordHasher interfaces to perform its operations.
type registerUserInteractor struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
}

// NewRegisterUserUseCase creates and returns a new instance of registerUserInteractor.
// Its dependencies (UserRepository and PasswordHasher) are injected from the outside.
func NewRegisterUserUseCase(
	userRepository UserRepository,
	passwordHasher PasswordHasher,
) RegisterUserUseCase {
	return &registerUserInteractor{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
	}
}

// Execute contains the core business logic for the user registration process.
func (r *registerUserInteractor) Execute(input RegisterUserInput) error {
	// 1. Basic input validation (more advanced validation can be added here or in a separate validator).
	if input.Email == "" || input.Password == "" {
		return fmt.Errorf("email and password cannot be empty")
	}

	// 2. Check if a user with the given email already exists.
	existingUser, err := r.userRepository.GetUserByEmail(input.Email)
	if err != nil && err != ErrUserNotFound {
		// Handle database errors other than "user not found" specifically.
		return fmt.Errorf("failed to check for existing user: %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user with this email already exists")
	}

	// 3. Hash the user's password using the PasswordHasher.
	hashedPassword, err := r.passwordHasher.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// 4. Create a new User entity with the hashed password.
	newUser := &entities.User{
		Email:    input.Email,
		Password: hashedPassword, // Store the hashed password
	}

	// 5. Persist the new user to the database via the UserRepository.
	err = r.userRepository.CreateUser(newUser)
	if err != nil {
		return fmt.Errorf("failed to save new user: %w", err)
	}

	return nil // Registration successful
}
