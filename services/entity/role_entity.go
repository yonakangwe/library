/**
 * @author Yohana Kangwe
 * @email yonakangwe@gmail.com
 * @create date 2025-01-05 13:13:37
 * @modify date 2025-01-05 13:13:37
 * @desc [description]
 */

package entity

import (
	"errors"
	"library/package/log"
	"time"
)

type Role struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	CreatedAt   time.Time
	UpdatedBy   int32
	UpdatedAt   time.Time
	DeletedBy   int32
	DeletedAt   time.Time
}

func NewRole(name, description string, createdBy int32) (*Role, error) {
	role := &Role{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}
	err := role.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Role entity %v", err)
		return nil, err
	}
	return role, nil
}

func (r *Role) ValidateCreate() error {
	if r.Name == "" {
		return errors.New("error validating Role entity, name field required")
	}
	if r.Description == "" {
		return errors.New("error validating Role entity, description field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Role entity, created_by field required")
	}
	return nil
}

func (r *Role) ValidateUpdate() error {
	if r.ID  <= 0 {
		return errors.New("error validating Role entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Role entity, name field required")
	}
	if r.Description == "" {
		return errors.New("error validating Role entity, description field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Role entity, updated_by field required")
	}
	return nil
}

type RoleFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Name      string
}
