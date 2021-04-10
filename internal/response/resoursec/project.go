package resources

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"time"
)

type ProjectResponse struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	CreatorId uint64 `json:"creator_id"`
	Creator *CreatorResponse `json:"creator"`
	Freelancer *FreelancerResponse `json:"freelancer"`
	FreelancerId *uint64 `json:"freelancer_id"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProjectResponse(project *models.Project) *ProjectResponse {
	return &ProjectResponse{
		Id:           project.Id,
		Name:         project.Name,
		Slug:         project.Slug,
		Description:  project.Description,
		IsActive:     project.IsActive,
		Freelancer:   NewFreelancerResponse(project.Freelancer),
		FreelancerId: project.FreelancerId,
		Price:        project.Price,
		CreatedAt:    project.CreatedAt,
	}
}

func NewProjectsResponse(projects []models.Project) *[]ProjectResponse {
	var res []ProjectResponse
	for _ , project := range projects {
		res = append(res , ProjectResponse{
			Id:           project.Id,
			Name:         project.Name,
			Slug:         project.Slug,
			Description:  project.Description,
			IsActive:     project.IsActive,
			CreatorId:    project.CreatorId,
			Creator:      NewCreatorResponse(project.Creator),
			Freelancer:   NewFreelancerResponse(project.Freelancer),
			FreelancerId: project.FreelancerId,
			Price:        project.Price,
			CreatedAt:    project.CreatedAt,
		})
	}

	return  &res
}