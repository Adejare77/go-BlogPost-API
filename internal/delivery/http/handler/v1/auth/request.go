package auth

type AuthRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3"`
}

type AuthRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3"`
}
