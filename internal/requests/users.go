package requests

type ProfileUpdate struct {
	LastName *string `json:"last_name"`
	FirstName *string `json:"first_name"`
	IsFreelancer *bool `json:"is_freelancer" binding:"required"`
	PhoneNumber *string `json:"phone_number"`
}
type ProfileAvatar struct {
	Avatar string `json:"avatar"`
}