package models

import (
	"time"
)

// Mkoa DataStructure
type Mkoa struct {
	ID        int32     `json:"id" validate:"numeric,required"`
	Name      string    `json:"name" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	Status    string    `json:"status"`
	CreatedBy int32     `json:"created_by" validate:"numeric,required"`
	UpdatedBy int32     `json:"updated_by" validate:"numeric,required"`
	DeletedBy int32     `json:"deleted_by" validate:"numeric,required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// MkoaFilter for list/pagination
type MkoaFilter struct {
	Filter
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}
