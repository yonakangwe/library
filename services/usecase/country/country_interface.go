package country

import (
	"library/services/entity"
)

type Reader interface {
	Get(id int32) (*entity.Country, error)
	List(filter *entity.CountryFilter) ([]*entity.Country, int32, error)
}

type Writer interface {
	Create(e *entity.Country) (int32, error)
	Update(e *entity.Country) error
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(name, isoCode string, phoneCode int16, createdBy int32) (int32, error)
	List(filter *entity.CountryFilter) ([]*entity.Country, int32, error)
	Get(id int32) (*entity.Country, error)
	Update(e *entity.Country) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}
