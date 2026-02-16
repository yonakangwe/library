package mkoa

import (
	"library/services/entity"
)

type Reader interface {
	Get(id int32) (*entity.Mkoa, error)
	GetByCode(code string) (*entity.Mkoa, error)
	List(filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error)
}

type Writer interface {
	Create(e *entity.Mkoa) (int32, error)
	Update(e *entity.Mkoa) error
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(name string, code string, createdBy int32) (int32, error)
	List(filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error)
	Get(id int32) (*entity.Mkoa, error)
	Update(e *entity.Mkoa) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}
