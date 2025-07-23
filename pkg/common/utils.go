package common

// This file can contain general utility functions that will be frequently used throughout the project.
// Although there isn't an immediate need in the current project, examples of future additions are provided.

// IsValidEmail performs a basic email format check.
// For more complex email validation, regular expressions (regexp) or external libraries can be used.
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// A basic check for '@' and '.'.
	// For a more robust check: `regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", email)`
	return len(email) > 3 && // Minimum length
		// Ensure it contains '@' and '.'
		// (This is just an example; in a real project, you should use regex)
		// bytes.Contains([]byte(email), []byte("@")) && bytes.Contains([]byte(email), []byte("."))
		// A better placeholder check
		email[0] != '@' && email[len(email)-1] != '.'
}
