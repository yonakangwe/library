package member

import (
	"library/services/entity"
)

type Reader interface {
	Get(id int32) (*entity.Member, error)
	// List(filter *entity.RoleFilter) ([]*entity.Role, int32, error)
}

type Writer interface {
	Create(e *entity.Member) (int32, error)
	Update(e *entity.Member) error
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(fullName string, phone string, email string, membershipNo string, createdBy int32) (int32, error)
	// List(filter *entity.RoleFilter) ([]*entity.Role, int32, error)
	Get(id int32) (*entity.Member, error)
	Update(e *entity.Member) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}
