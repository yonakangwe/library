package member

import (
	"library/services/entity"
)

type Reader interface {
	// Get(id int32) (*entity.Member, error)
	// List(filter *entity.MemberFilter) ([]*entity.Member, int32, error)
}

type Writer interface {
	Create(e *entity.Member) (int32, error)
	// Update(e *entity.Member) error
	// SoftDelete(id, deletedBy int32) error
	// HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(fullName string, email string, createdBy int32) (int32, error)
	// List(filter *entity.MemberFilter) ([]*entity.Member, int32, error)
	// Get(id int32) (*entity.Member, error)
	// Update(e *entity.Member) (int32, error)
	// SoftDelete(id, deletedBy int32) error
	// HardDelete(id int32) error
}
