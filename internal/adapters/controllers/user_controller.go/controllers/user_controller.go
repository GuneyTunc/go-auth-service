package controllers

import (
	"encoding/json"
	"net/http"

	// Importing use cases, because controllers depend on use cases
	"LoginMechanism/internal/application/usecases"
)

// UserController is used to handle HTTP requests.
// It takes RegisterUser and LoginUser use cases as its dependencies.
type UserController struct {
	registerUser usecases.RegisterUserUseCase
	loginUser    usecases.LoginUserUseCase
}

// NewUserController creates and returns a new instance of UserController.
// Its dependencies (use cases) are injected from the outside.
func NewUserController(
	registerUser usecases.RegisterUserUseCase,
	loginUser usecases.LoginUserUseCase,
) *UserController {
	return &UserController{
		registerUser: registerUser,
		loginUser:    loginUser,
	}
}

// RegisterUserRequest defines the structure of the incoming request body for the registration endpoint.
type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser is the HTTP handler that processes new user registration requests.
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Sending the input model to the use case
	input := usecases.RegisterUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// Executing the use case
	err = c.registerUser.Execute(input)
	if err != nil {
		// An HTTP status code can be returned based on the error type
		http.Error(w, "Registration failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUserRequest defines the structure of the incoming request body for the login endpoint.
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUser is the HTTP handler that processes user login requests.
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Sending the input model to the use case
	input := usecases.LoginUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// Executing the use case; if successful, it returns a JWT token
	output, err := c.loginUser.Execute(input)
	if err != nil {
		// An HTTP status code can be returned based on the error type (e.g., StatusUnauthorized)
		http.Error(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": output.Token})
}
