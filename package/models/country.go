package models

import (
	"time"
)

// Country DataStructure
type Country struct {
	ID        int32     `json:"id" validate:"numeric,required"`
	Name      string    `json:"name" validate:"required"`
	IsoCode   string    `json:"iso_code" validate:"required"`
	PhoneCode int16     `json:"phone_code" validate:"numeric,required"`
	IsActive  bool      `json:"is_active"`
	CreatedBy int32     `json:"created_by" validate:"numeric,required"`
	UpdatedBy int32     `json:"updated_by" validate:"numeric,required"`
	DeletedBy int32     `json:"deleted_by" validate:"numeric,required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CountryFilter struct {
	Filter
	Name string `json:"name"`
}
