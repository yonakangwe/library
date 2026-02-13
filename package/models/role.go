package models

import (
	"time"
)

// Role DataStructure
type Role struct {
	ID          int32     `json:"id" validate:"numeric,required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedBy   int32     `json:"created_by" validate:"numeric,required"`
	UpdatedBy   int32     `json:"updated_by" validate:"numeric,required"`
	DeletedBy   int32     `json:"deleted_by" validate:"numeric,required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type RoleFilter struct {
	Filter
	Name int32 `json:"name"`
}
