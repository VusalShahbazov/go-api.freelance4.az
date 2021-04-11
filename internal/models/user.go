package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id uint64 `json:"id"`
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email string `json:"email"`
	Password string `json:"-"`
	PhoneNumber *string `json:"phone_number"`
	Avatar *string `json:"avatar"`
	IsFreelancer bool `json:"is_freelancer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}