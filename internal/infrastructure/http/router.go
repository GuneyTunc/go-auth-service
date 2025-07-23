package http

import (
	"log"
	"net/http"
)

// NewRouter, HTTP isteklerini yönlendirmek için bir http.ServeMux instance'ı döndürür.
// Bu, Go'nun standart kütüphanesindeki basit bir HTTP router'ıdır.
func NewRouter() *http.ServeMux {
	// http.NewServeMux, HTTP isteklerini URL yollarına göre handler'lara yönlendiren bir multiplexer'dır.
	mux := http.NewServeMux()

	// Örneğin, tüm istekler için genel bir loglama middleware'i eklenebilir.
	// Ancak, Clean Architecture'da middleware'ler genellikle ana router'ın üstünde,
	// yani main.go'da veya burada daha karmaşık bir wrapper ile tanımlanır.
	// Şimdilik basit tutalım.

	log.Println("HTTP router initialized.")
	return mux
}

// HandlerFuncWrapper, http.HandlerFunc'a uygun bir adapter oluşturur.
// Bu, Controller metodlarını doğrudan http.HandleFunc'a bağlamak için bir yardımcı fonksiyondur.
// İleride daha gelişmiş middleware'ler veya error handling için genişletilebilir.
type HandlerFuncWrapper func(http.ResponseWriter, *http.Request)

// Methods is a helper to chain methods, similar to some web frameworks.
// This is a simplified example to make routing definitions cleaner.
// For more robust routing, consider a dedicated router library (Gorilla Mux, Gin, Echo).
func (h HandlerFuncWrapper) Methods(methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := false
		for _, m := range methods {
			if r.Method == m {
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}
