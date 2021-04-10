package resources

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
)

type CreatorResponse struct {
	Id uint64 `json:"id"`
	FullName string `json:"full_name"`
}

type FreelancerResponse struct {
	Id uint64 `json:"id"`
	FullName string `json:"full_name"`
}

func NewCreatorResponse(user *models.User) *CreatorResponse {
	if user == nil {
		return  nil
	}
	return &CreatorResponse{
		Id:       user.Id,
		FullName: user.FirstName + " " +user.LastName,
	}
}

func NewFreelancerResponse(user *models.User) *FreelancerResponse {
	if user == nil {
		return  nil
	}
	return &FreelancerResponse{
		Id:       user.Id,
		FullName: user.FirstName + " " +user.LastName,
	}
}