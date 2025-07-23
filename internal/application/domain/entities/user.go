package entities

// User represents the core business entity of the application.
// This entity is independent of external technologies like databases or HTTP.
type User struct {
	ID       int    // The user's unique identifier, typically assigned by the database
	Email    string // The user's email address (must be unique)
	Password string // The user's hashed password
}

// NewUser is a helper function to create a new User entity.
// This is generally the preferred method for creating entities in the domain layer.
func NewUser(id int, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}
