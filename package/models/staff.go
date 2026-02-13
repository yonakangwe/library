package models

import (
	"time"
)

// Role DataStructure
type Staff struct {
	ID           int32     `json:"id" validate:"numeric,required"`
	FullName     string    `json:"name" validate:"required"`
	Email        string    `json:"description" validate:"required"`
	Phone        string    `json:"phone" validate:"required"`
	Username     string    `json:"username" validate:"required"`
	PasswordHash string    `json:"password_hash" validate:"required"`
	CreatedBy    int32     `json:"created_by" validate:"numeric,required"`
	UpdatedBy    int32     `json:"updated_by" validate:"numeric,required"`
	DeletedBy    int32     `json:"deleted_by" validate:"numeric,required"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
