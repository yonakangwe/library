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

func (m *Member) ValidateCreate() error {
	if m.FullName == "" {
		return errors.New("error validating Members entity, FullName field required")
	}
	if m.Phone == "" {
		return errors.New("error validating Members entity, Phone field required")
	}
	if m.Email == "" {
		return errors.New("error validating Members entity, Email field required")
	}
	if m.MembershipNo <= 0 {
		return errors.New("error validating Members entity, MembershipNo field required")
	}
	if m.CreatedBy <= 0 {
		return errors.New("error validating Members entity, CreatedBy field required")
	}
	return nil
}
