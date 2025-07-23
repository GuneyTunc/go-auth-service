package repositories

import (
	"database/sql"
	"fmt"

	// Domain entity'yi dahil ediyoruz
	"LoginMechanism/internal/application/domain/entities"
	// Use case'lerin bağımlı olduğu Repository arayüzünü dahil ediyoruz
	"LoginMechanism/internal/application/usecases"
)

// SQLUserRepository, UserRepository arayüzünün bir SQL Server implementasyonudur.
// Bir *sql.DB bağlantısına bağımlıdır.
type SQLUserRepository struct {
	db *sql.DB
}

// NewSQLUserRepository yeni bir SQLUserRepository instance'ı oluşturur.
// Veritabanı bağlantısı dışarıdan enjekte edilir.
func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// CreateUser, yeni bir User'ı veritabanına ekler.
// Bu metot, usecases.UserRepository arayüzünün CreateUser metodunu uygular.
func (r *SQLUserRepository) CreateUser(user *entities.User) error {
	// SQL Server'a özgü parametre (@p1, @p2) kullanımı.
	// Güvenlik için parametreli sorgular kullanılır.
	query := "INSERT INTO users (email, password) VALUES (@p1, @p2);"

	_, err := r.db.Exec(query, user.Email, user.Password)
	if err != nil {
		// E-posta benzersizliği hatası gibi spesifik hatalar burada yakalanabilir
		// ve daha genel bir domain hatasına dönüştürülebilir.
		return fmt.Errorf("failed to create user in database: %w", err)
	}
	return nil
}

// GetUserByEmail, verilen e-posta adresine sahip User'ı veritabanından getirir.
// Bu metot, usecases.UserRepository arayüzünün GetUserByEmail metodunu uygular.
func (r *SQLUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = @p1;"

	user := &entities.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		// Kullanıcı bulunamazsa özel bir hata döndürüyoruz.
		// Bu hatayı use case katmanında kontrol edebiliriz.
		return nil, usecases.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email from database: %w", err)
	}

	return user, nil
}
