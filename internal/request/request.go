package request

type AuthRequest struct {
	Email    string `json:"email" validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=5,max=255"`
}
