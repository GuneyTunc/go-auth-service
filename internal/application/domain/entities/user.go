package entities

// User, uygulamanın çekirdek iş varlığını temsil eder.
// Bu entity, veritabanı veya HTTP gibi dış teknolojilerden bağımsızdır.
type User struct {
	ID       int    // Kullanıcının benzersiz kimliği, genellikle veritabanı tarafından atanır
	Email    string // Kullanıcının e-posta adresi (benzersiz olmalı)
	Password string // Kullanıcının hashlenmiş şifresi
}

// NewUser, yeni bir User entity'si oluşturmak için bir yardımcı fonksiyondur.
// Bu, genellikle domain katmanında entity'leri oluşturmak için tercih edilen bir yöntemdir.
func NewUser(id int, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

// Örneğin, User entity'sine özgü basit bir doğrulama veya mantık burada yer alabilir.
// func (u *User) IsPasswordValid(hashedPassword string) bool {
//     // Bu kontrol aslında PasswordHasher servisinde yapılmalı,
//     // ancak domain entity'sine ait basit bir kural örneği olabilir.
//     return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password)) == nil
// }

// func (u *User) ChangeEmail(newEmail string) error {
//     // E-posta değiştirme gibi bir domain davranışı burada olabilir.
//     // Ancak bu projede doğrudan repository tarafından yönetiliyor.
//     u.Email = newEmail
//     return nil
// }
