package entity

import (
	"errors"
	"library/package/log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Staff struct {
	ID           int32
	FullName     string
	Email        string
	Phone        string
	Username     string
	PasswordHash string
	IsActive     bool
	CreatedBy    int32
	CreatedAt    time.Time
	UpdatedBy    int32
	UpdatedAt    time.Time
	DeletedBy    int32
	DeletedAt    time.Time
}

func NewStaff(fullname, email, phone, username, passwordHash string, createdBy int32) (*Staff, error) {
	staff := &Staff{
		FullName:     fullname,
		Email:        email,
		Phone:        phone,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedBy:    createdBy,
	}

	// Hashing password
	if err := staff.EncryptPassword(); err != nil {
		log.Errorf("error encrypting password %v", err)
		return nil, err
	}
	err := staff.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Staff entity %v", err)
		return nil, err
	}
	return staff, nil
}

func (r *Staff) ValidateCreate() error {
	if r.FullName == "" {
		return errors.New("error validating Staff entity, name field required")
	}
	if r.Email == "" {
		return errors.New("error validating Staff entity, description field required")
	}
	if r.Phone == "" {
		return errors.New("error validating Staff entity, phone field required")
	}
	if r.Username == "" {
		return errors.New("error validating Staff entity, username field required")
	}
	if r.PasswordHash == "" {
		return errors.New("error validating Staff entity, password_hash field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Staff entity, created_by field required")
	}
	return nil
}

func (r *Staff) ValidateUpdate() error {
	if r.FullName == "" {
		return errors.New("error validating Staff entity, name field required")
	}
	if r.Email == "" {
		return errors.New("error validating Staff entity, description field required")
	}
	if r.Phone == "" {
		return errors.New("error validating Staff entity, phone field required")
	}
	if r.Username == "" {
		return errors.New("error validating Staff entity, username field required")
	}
	if r.PasswordHash == "" {
		return errors.New("error validating Staff entity, password_hash field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Staff entity, updated_by field required")
	}
	return nil
}

func (r *Staff) EncryptPassword() error {
	if r.PasswordHash == "" {
		return errors.New("password is required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(r.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	r.PasswordHash = string(hashed)
	r.PasswordHash = "" // clear plain password for security
	return nil
}
