package entity

import (
	"errors"
	"library/package/log"
	"time"
)

type Member struct {
	ID           int32
	FullName     string
	Phone        string
	Email        string
	MembershipNo int32
	CreatedBy    int32
	CreatedAt    time.Time
	UpdatedBy    int32
	UpdatedAt    time.Time
	DeletedBy    int32
	DeletedAt    time.Time
}

func NewMember(fullName, phone, email string, membershipNo, createdBy int32) (*Member, error) {
	member := &Member{
		FullName:     fullName,
		Phone:        phone,
		Email:        email,
		MembershipNo: membershipNo,
		CreatedBy:    createdBy,
	}

	err := member.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Member entity %v", err)
		return nil, err
	}
	return member, nil
}

func (r *Member) ValidateCreate() error {
	if r.FullName == "" {
		return errors.New("error validating Members entity, FullName field required")
	}
	if r.Phone == "" {
		return errors.New("error validating Members entity, description field required")
	}
	if r.Email == "" {
		return errors.New("error validating Members entity, Phone field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Members entity, created_by field required")
	}
	return nil
}
