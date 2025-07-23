package repositories

import (
	"database/sql"
	"fmt"

	// Include the domain entity
	"LoginMechanism/internal/application/domain/entities"
	// Include the Repository interface that use cases depend on
	"LoginMechanism/internal/application/usecases"
)

// SQLUserRepository is a SQL Server implementation of the UserRepository interface.
// It depends on an *sql.DB database connection.
type SQLUserRepository struct {
	db *sql.DB
}

// NewSQLUserRepository creates and returns a new instance of SQLUserRepository.
// The database connection is injected from the outside.
func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// CreateUser adds a new User to the database.
// This method implements the CreateUser method of the usecases.UserRepository interface.
func (r *SQLUserRepository) CreateUser(user *entities.User) error {
	// SQL Server-specific parameter usage (@p1, @p2).
	// Parameterized queries are used for security.
	query := "INSERT INTO users (email, password) VALUES (@p1, @p2);"

	_, err := r.db.Exec(query, user.Email, user.Password)
	if err != nil {
		// Specific errors like email uniqueness violation can be caught here
		// and converted into a more general domain error.
		return fmt.Errorf("failed to create user in database: %w", err)
	}
	return nil
}

// GetUserByEmail retrieves a User from the database with the given email address.
// This method implements the GetUserByEmail method of the usecases.UserRepository interface.
func (r *SQLUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = @p1;"

	user := &entities.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		// We return a specific error if the user is not found.
		// This error can be handled in the use case layer.
		return nil, usecases.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email from database: %w", err)
	}

	return user, nil
}
