package role

import (
	"library/services/entity"
)

type Reader interface {
	//Get(id int32) (*entity.Role, error)
	//List(filter *entity.RoleFilter) ([]*entity.Role, int32, error)
}

type Writer interface {
	Create(e *entity.Role) (int32, error)
	//Update(e *entity.Role) error
	//SoftDelete(id, deletedBy int32) error
	//HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateRole(name string, description string, createdBy int32) (int32, error)
	// List(filter *entity.RoleFilter) ([]*entity.Role, int32, error)
	// Get(id int32) (*entity.Role, error)
	// Update(e *entity.Role) (int32, error)
	// SoftDelete(id, deletedBy int32) error
	// HardDelete(id int32) error
}
