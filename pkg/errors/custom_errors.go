package errors

import "fmt"

// ValidationError represents errors that occur during validation.
// It can contain multiple validation error messages.
type ValidationError struct {
	Messages map[string]string // Field name: Error message
}

// Error returns a string representation of the ValidationError.
func (e *ValidationError) Error() string {
	msg := "validation error(s):"
	for field, errMsg := range e.Messages {
		msg += fmt.Sprintf(" %s: %s;", field, errMsg)
	}
	return msg
}

// NewValidationError creates and returns a new instance of ValidationError.
func NewValidationError(messages map[string]string) *ValidationError {
	return &ValidationError{
		Messages: messages,
	}
}

// IsValidationError checks if the given error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// NotFoundError represents an error indicating that a resource could not be found.
type NotFoundError struct {
	ResourceName string // The name of the resource that was not found (e.g., "User", "Product")
	Identifier   string // The identifier of the resource (e.g., "email@example.com", "123")
}

// Error returns a string representation of the NotFoundError.
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with identifier '%s' not found", e.ResourceName, e.Identifier)
}

// NewNotFoundError creates and returns a new instance of NotFoundError.
func NewNotFoundError(resourceName, identifier string) *NotFoundError {
	return &NotFoundError{
		ResourceName: resourceName,
		Identifier:   identifier,
	}
}

// IsNotFoundError checks if the given error is a NotFoundError.
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// ConflictError represents an error indicating a conflict in an operation.
// It is typically used when unique constraints are violated (e.g., an email that already exists).
type ConflictError struct {
	Message string
}

// Error returns a string representation of the ConflictError.
func (e *ConflictError) Error() string {
	return e.Message
}

// NewConflictError creates and returns a new instance of ConflictError.
func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}

// IsConflictError checks if the given error is a ConflictError.
func IsConflictError(err error) bool {
	_, ok := err.(*ConflictError)
	return ok
}

// UnauthorizedError represents a failed authentication (invalid credentials).
type UnauthorizedError struct {
	Message string
}

// Error returns a string representation of the UnauthorizedError.
func (e *UnauthorizedError) Error() string {
	return e.Message
}

// NewUnauthorizedError creates and returns a new instance of UnauthorizedError.
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
	}
}

// IsUnauthorizedError checks if the given error is an UnauthorizedError.
func IsUnauthorizedError(err error) bool {
	_, ok := err.(*UnauthorizedError)
	return ok
}

// InternalServerError represents unexpected or unhandled internal system errors.
type InternalServerError struct {
	Message string
	Err     error // Optional: to wrap the original error
}

// Error returns a string representation of the InternalServerError.
func (e *InternalServerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("internal server error: %s (details: %v)", e.Message, e.Err)
	}
	return fmt.Sprintf("internal server error: %s", e.Message)
}

// NewInternalServerError creates and returns a new instance of InternalServerError.
func NewInternalServerError(message string, err error) *InternalServerError {
	return &InternalServerError{
		Message: message,
		Err:     err,
	}
}

// IsInternalServerError checks if the given error is an InternalServerError.
func IsInternalServerError(err error) bool {
	_, ok := err.(*InternalServerError)
	return ok
}
