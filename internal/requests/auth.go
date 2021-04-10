package requests

type SignUpRequest struct {
	LastName string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResetPassword struct {
	NewPassword string `json:"new_password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`

}