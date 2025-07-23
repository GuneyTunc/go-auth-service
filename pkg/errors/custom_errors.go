package errors

import "fmt"

// ValidationError, geçerlilik kontrolü (validation) sırasında oluşan hataları temsil eder.
// Birden fazla doğrulama hatası içerebilir.
type ValidationError struct {
	Messages map[string]string // Alan adı: Hata mesajı
}

func (e *ValidationError) Error() string {
	msg := "validation error(s):"
	for field, errMsg := range e.Messages {
		msg += fmt.Sprintf(" %s: %s;", field, errMsg)
	}
	return msg
}

// NewValidationError, yeni bir ValidationError instance'ı oluşturur.
func NewValidationError(messages map[string]string) *ValidationError {
	return &ValidationError{
		Messages: messages,
	}
}

// IsValidationError, verilen hatanın bir ValidationError olup olmadığını kontrol eder.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// NotFoundError, bir kaynağın (resource) bulunamadığını belirten hatayı temsil eder.
type NotFoundError struct {
	ResourceName string // Hangi kaynağın bulunamadığı (örn. "User", "Product")
	Identifier   string // Kaynağın tanımlayıcısı (örn. "email@example.com", "123")
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with identifier '%s' not found", e.ResourceName, e.Identifier)
}

// NewNotFoundError, yeni bir NotFoundError instance'ı oluşturur.
func NewNotFoundError(resourceName, identifier string) *NotFoundError {
	return &NotFoundError{
		ResourceName: resourceName,
		Identifier:   identifier,
	}
}

// IsNotFoundError, verilen hatanın bir NotFoundError olup olmadığını kontrol eder.
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// ConflictError, bir işlemde çakışma (conflict) olduğunu belirten hatayı temsil eder.
// Genellikle benzersiz kısıtlamaların ihlal edilmesi (örn. zaten var olan bir e-posta) durumunda kullanılır.
type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

// NewConflictError, yeni bir ConflictError instance'ı oluşturur.
func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}

// IsConflictError, verilen hatanın bir ConflictError olup olmadığını kontrol eder.
func IsConflictError(err error) bool {
	_, ok := err.(*ConflictError)
	return ok
}

// UnauthorizedError, kimlik doğrulama başarısızlığını (geçersiz kimlik bilgileri) temsil eder.
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

// NewUnauthorizedError, yeni bir UnauthorizedError instance'ı oluşturur.
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
	}
}

// IsUnauthorizedError, verilen hatanın bir UnauthorizedError olup olmadığını kontrol eder.
func IsUnauthorizedError(err error) bool {
	_, ok := err.(*UnauthorizedError)
	return ok
}

// InternalServerError, beklenmedik veya işlenemeyen iç sistem hatalarını temsil eder.
type InternalServerError struct {
	Message string
	Err     error // Orijinal hatayı sarmalamak için (opsiyonel)
}

func (e *InternalServerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("internal server error: %s (details: %v)", e.Message, e.Err)
	}
	return fmt.Sprintf("internal server error: %s", e.Message)
}

// NewInternalServerError, yeni bir InternalServerError instance'ı oluşturur.
func NewInternalServerError(message string, err error) *InternalServerError {
	return &InternalServerError{
		Message: message,
		Err:     err,
	}
}

// IsInternalServerError, verilen hatanın bir InternalServerError olup olmadığını kontrol eder.
func IsInternalServerError(err error) bool {
	_, ok := err.(*InternalServerError)
	return ok
}
