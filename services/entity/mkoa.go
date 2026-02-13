package entity

import (
	"errors"
	"time"
)

const (
	MkoaStatusActive   = "active"
	MkoaStatusInactive = "inactive"
)

type Mkoa struct {
	ID        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Code      string     `db:"code" json:"code"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	CreatedBy int64      `db:"created_by" json:"created_by"`
	UpdatedBy int64      `db:"updated_by" json:"updated_by"`
	DeletedBy int64      `db:"deleted_by" json:"deleted_by"`
}

func NewMkoa(name, code string, createdBy int64) (*Mkoa, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if code == "" {
		return nil, errors.New("code is required")
	}
	if createdBy <= 0 {
		return nil, errors.New("created_by is required")
	}

	now := time.Now()

	return &Mkoa{
		Name:      name,
		Code:      code,
		Status:    MkoaStatusActive,
		CreatedAt: now,
		CreatedBy: createdBy,
	}, nil
}

func (m *Mkoa) ValidateCreate() error {
	if m.Name == "" {
		return errors.New("error validating Mkoa entity, name field required")
	}
	if m.Code == "" {
		return errors.New("error validating Mkoa entity, code field required")
	}
	if m.CreatedBy <= 0 {
		return errors.New("error validating Mkoa entity, created_by field required")
	}
	if m.Status == "" {
		m.Status = MkoaStatusActive
	}
	return nil
}

func (m *Mkoa) ValidateUpdate() error {
	if m.ID <= 0 {
		return errors.New("error validating Mkoa entity, ID field required")
	}
	if m.Name == "" {
		return errors.New("error validating Mkoa entity, name field required")
	}
	if m.Code == "" {
		return errors.New("error validating Mkoa entity, code field required")
	}
	if m.UpdatedBy <= 0 {
		return errors.New("error validating Mkoa entity, updated_by field required")
	}
	return nil
}

func (m *Mkoa) ValidateDelete() error {
	if m.ID <= 0 {
		return errors.New("error validating Mkoa entity, ID field required")
	}
	if m.DeletedBy <= 0 {
		return errors.New("error validating Mkoa entity, deleted_by field required")
	}
	return nil
}

type MkoaFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Name      string
	Code      string
	Status    string
}
