package main

import (
	"fmt"
	"log"
	"net/http"

	// Adapters
	"LoginMechanism/internal/adapters/auth"
	"LoginMechanism/internal/adapters/controllers/user_controller.go/controllers"
	"LoginMechanism/internal/adapters/repositories"

	// Application (Use Cases)
	"LoginMechanism/internal/application/usecases"

	// Infrastructure
	"LoginMechanism/internal/infrastructure/config"
	database "LoginMechanism/internal/infrastructure/databese"
	myhttp "LoginMechanism/internal/infrastructure/http"

	// Domain services (if any are externalized for DI)
	"LoginMechanism/internal/application/domain/services" // For PasswordHasher implementation

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// 1. Load Configuration (from .env file)
	config.LoadEnv("config.env")

	connString := config.GetEnv("DB_CONNECTION_STRING")
	jwtSecretKey := config.GetEnv("JWT_SECRET_KEY")

	// 2. Initialize Infrastructure Components
	// Database connection
	db, err := database.NewSQLServerDB(connString) // This function will handle opening/pinging
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close() // Ensure DB connection is closed when main exits
	fmt.Println("Successfully connected to the database! 🚀")

	// 3. Initialize Domain Services (implementations of interfaces defined in Application/Domain)
	passwordHasher := services.NewBcryptPasswordHasher() // bcrypt'i kullanan somut PasswordHasher implementasyonu

	// 4. Initialize Adapters (Repositories & Auth Providers)
	userRepo := repositories.NewSQLUserRepository(db) // SQL tabanlı UserRepository implementasyonu
	tokenProvider := auth.NewJWTTokenProvider([]byte(jwtSecretKey))

	// 5. Initialize Application Use Cases (Core Business Logic)
	registerUserUseCase := usecases.NewRegisterUserUseCase(userRepo, passwordHasher)
	loginUserUseCase := usecases.NewLoginUserUseCase(userRepo, passwordHasher, tokenProvider)

	// 6. Initialize Controllers (HTTP Handlers)
	userController := controllers.NewUserController(registerUserUseCase, loginUserUseCase)

	// 7. Setup HTTP Router and Routes
	r := myhttp.NewRouter() // Create a new router instance
	r.HandleFunc("/register", myhttp.HandlerFuncWrapper(userController.RegisterUser).Methods("POST"))
	r.HandleFunc("/login", myhttp.HandlerFuncWrapper(userController.LoginUser).Methods("POST"))

	// 8. Start the HTTP Server
	fmt.Println("Server is running on port 8080... 🌐")
	log.Fatal(http.ListenAndServe(":8080", r)) // Pass the router to ListenAndServe
}
