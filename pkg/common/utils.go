package common

// Bu dosyada projede sıkça kullanılacak genel yardımcı (utility) fonksiyonlar bulunabilir.
// Mevcut projede acil bir ihtiyaç olmasa da, ileride eklenebilecek örneklere yer verilmiştir.

// IsValidEmail basit bir e-posta formatı kontrolü yapar.
// Daha karmaşık e-posta validasyonları için regexp veya harici kütüphaneler kullanılabilir.
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// Basit bir '@' ve '.' kontrolü
	// Daha sağlam bir kontrol için: `regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", email)`
	return len(email) > 3 && // Minimum uzunluk
		// '@' ve '.' içerdiğinden emin ol
		// (Bu sadece bir örnek, gerçek projede regex kullanmalısınız)
		// bytes.Contains([]byte(email), []byte("@")) && bytes.Contains([]byte(email), []byte("."))
		// Daha iyi bir yer tutucu kontrolü
		email[0] != '@' && email[len(email)-1] != '.'
}

// GenerateRandomString belirli uzunlukta rastgele bir string oluşturur.
// Genellikle ID'ler, tokenlar veya geçici şifreler için kullanılır.
// Bu proje için doğrudan gerekli olmasa da, sık karşılaşılan bir yardımcı fonksiyondur.
// import "crypto/rand"
// import "encoding/base64"
//
// func GenerateRandomString(length int) (string, error) {
// 	b := make([]byte, length)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(b)[:length], nil
// }

// ContainsStringSlice bir string slice'ın belirli bir string içerip içermediğini kontrol eder.
// func ContainsStringSlice(slice []string, val string) bool {
// 	for _, item := range slice {
// 		if item == val {
// 			return true
// 		}
// 	}
// 	return false
// }
