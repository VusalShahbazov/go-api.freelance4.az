package requests

type ProjectFilter struct {
	Pagination
	IsActive *bool `uri:"is_active" binding:"omitempty"`
}