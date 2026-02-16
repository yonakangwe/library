package entity

import (
	"errors"
	"library/package/log"
	"time"
)

type Member struct {
	ID        int32
	FullName  string
	Email     string
	CreatedBy int32
	CreatedAt time.Time
	UpdatedBy int32
	UpdatedAt time.Time
	DeletedBy int32
	DeletedAt time.Time
}

func NewMember(fullName, email string, createdBy int32) (*Member, error) {
	member := &Member{
		FullName:  fullName,
		Email:     email,
		CreatedBy: createdBy,
	}

	err := member.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Member entity %v", err)
		return nil, err
	}
	return member, nil
}

func (m *Member) ValidateCreate() error {
	if m.FullName == "" {
		return errors.New("error validating Members entity, FullName field required")
	}
	if m.Email == "" {
		return errors.New("error validating Members entity, Email field required")
	}
	if m.CreatedBy <= 0 {
		return errors.New("error validating Members entity, CreatedBy field required")
	}
	return nil
}

func (r *Member) ValidateUpdate() error {
	if r.ID <= 0 {
		return errors.New("error validating Member entity, ID field required")
	}
	if r.FullName == "" {
		return errors.New("error validating Member entity, fullname field required")
	}
	if r.Email == "" {
		return errors.New("error validating Member entity, email field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Member entity, updated_by field required")
	}
	return nil
}

type MemberFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	FullName  string
}
