package user

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password" binding:"omitempty,min=3"`
}
