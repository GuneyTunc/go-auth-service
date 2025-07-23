package controllers

import (
	"encoding/json"
	"net/http"

	// Use cases'leri dahil ediyoruz, çünkü controller'lar use case'lere bağımlıdır
	"LoginMechanism/internal/application/usecases"
)

// UserController HTTP isteklerini işlemek için kullanılır.
// Bağımlılıkları olarak RegisterUser ve LoginUser use case'lerini alır.
type UserController struct {
	registerUser usecases.RegisterUserUseCase
	loginUser    usecases.LoginUserUseCase
}

// NewUserController yeni bir UserController instance'ı oluşturur.
// Bağımlılıkları (use case'ler) dışarıdan enjekte edilir.
func NewUserController(
	registerUser usecases.RegisterUserUseCase,
	loginUser usecases.LoginUserUseCase,
) *UserController {
	return &UserController{
		registerUser: registerUser,
		loginUser:    loginUser,
	}
}

// RegisterUserRequest, kayıt endpoint'i için gelen isteğin yapısını tanımlar.
type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterUser, yeni kullanıcı kayıt isteklerini işleyen HTTP handler'ıdır.
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

	// Use case'e input modelini gönderiyoruz
	input := usecases.RegisterUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// Use case'i çalıştırıyoruz
	err = c.registerUser.Execute(input)
	if err != nil {
		// Hata türüne göre HTTP durum kodu döndürülebilir
		http.Error(w, "Registration failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUserRequest, giriş endpoint'i için gelen isteğin yapısını tanımlar.
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUser, kullanıcı giriş isteklerini işleyen HTTP handler'ıdır.
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

	// Use case'e input modelini gönderiyoruz
	input := usecases.LoginUserInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// Use case'i çalıştırıyoruz, başarılı olursa JWT token'ı döner
	output, err := c.loginUser.Execute(input)
	if err != nil {
		// Hata türüne göre HTTP durum kodu döndürülebilir (örn. StatusUnauthorized)
		http.Error(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": output.Token})
}
