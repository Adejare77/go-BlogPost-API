package user

type UserCreateRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password" binding:"omitempty,min=3"`
}
