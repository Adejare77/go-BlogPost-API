package user

type UserPathRequest struct {
	UserID int `uri:"user_id" binding:"required"`
}
