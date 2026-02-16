package entity

import (
	"errors"
	"fmt"
	"time"
)

const (
	MkoaStatusActive   = "active"
	MkoaStatusInactive = "inactive"
)

// ErrMkoaInvalid is a sentinel error for all Mkoa validation failures.
// Callers can use errors.Is(err, ErrMkoaInvalid) and still inspect the message for details.
var ErrMkoaInvalid = errors.New("invalid mkoa")

type Mkoa struct {
	ID        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Code      string     `db:"code" json:"code"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	CreatedBy *int64     `db:"created_by" json:"created_by,omitempty"`
	UpdatedBy *int64     `db:"updated_by" json:"updated_by,omitempty"`
	DeletedBy *int64     `db:"deleted_by" json:"deleted_by,omitempty"`
}

// Int64PtrVal returns the value or 0 if nil (for nullable DB columns).
func Int64PtrVal(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

// IsActive is a small convenience helper for checking the current status.
func (m *Mkoa) IsActive() bool {
	return m.Status == MkoaStatusActive
}

func NewMkoa(name, code string, createdBy int64) (*Mkoa, error) {
	if name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrMkoaInvalid)
	}
	if code == "" {
		return nil, fmt.Errorf("%w: code is required", ErrMkoaInvalid)
	}
	if createdBy <= 0 {
		return nil, fmt.Errorf("%w: created_by is required", ErrMkoaInvalid)
	}

	now := time.Now()

	cb := createdBy
	return &Mkoa{
		Name:      name,
		Code:      code,
		Status:    MkoaStatusActive,
		CreatedAt: now,
		CreatedBy: &cb,
	}, nil
}

func (m *Mkoa) ValidateCreate() error {
	if m.Name == "" {
		return fmt.Errorf("%w: name field required", ErrMkoaInvalid)
	}
	if m.Code == "" {
		return fmt.Errorf("%w: code field required", ErrMkoaInvalid)
	}
	if m.CreatedBy == nil || *m.CreatedBy <= 0 {
		return fmt.Errorf("%w: created_by field required", ErrMkoaInvalid)
	}
	if m.Status == "" {
		m.Status = MkoaStatusActive
	}
	return nil
}

func (m *Mkoa) ValidateUpdate() error {
	if m.ID <= 0 {
		return fmt.Errorf("%w: ID field required", ErrMkoaInvalid)
	}
	if m.Name == "" {
		return fmt.Errorf("%w: name field required", ErrMkoaInvalid)
	}
	if m.Code == "" {
		return fmt.Errorf("%w: code field required", ErrMkoaInvalid)
	}
	if m.UpdatedBy == nil || *m.UpdatedBy <= 0 {
		return fmt.Errorf("%w: updated_by field required", ErrMkoaInvalid)
	}
	return nil
}

func (m *Mkoa) ValidateDelete() error {
	if m.ID <= 0 {
		return fmt.Errorf("%w: ID field required", ErrMkoaInvalid)
	}
	if m.DeletedBy == nil || *m.DeletedBy <= 0 {
		return fmt.Errorf("%w: deleted_by field required", ErrMkoaInvalid)
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
