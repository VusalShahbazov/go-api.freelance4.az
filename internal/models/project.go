package models

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	CreatorId uint64 `json:"creator_id"`
	Creator *User `json:"creator" gorm:"foreignKey:CreatorId"`
	Freelancer *User `json:"freelancer" gorm:"foreignKey:FreelancerId"`
	FreelancerId *uint64 `json:"freelancer_id"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
