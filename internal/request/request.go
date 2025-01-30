package request

// AuthRequest Тіло запита авторизації.
type AuthRequest struct {
	Email    string `json:"email" validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=5,max=255"`
}

// RetrospectiveRequest Тіло запита на перевірку токена.
type RetrospectiveRequest struct {
	Jwt string `form:"jwt"`
}
