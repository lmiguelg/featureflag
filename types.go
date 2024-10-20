package main

import (
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsActive    bool      `json:"isActive"`
}

type CreateProjectRequest struct {
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

func NewProject(desc string, isActive bool) *Project {
	return &Project{
		Description: desc,
		IsActive:    isActive,
		CreatedAt:   time.Now().UTC(),
	}
}
