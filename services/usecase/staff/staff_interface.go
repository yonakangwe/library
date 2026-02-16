package staff

import (
	"library/services/entity"
)

type Reader interface {
}

type Writer interface {
	Create(e *entity.Staff) (int32, error)
	Update(e *entity.Staff) (int32, error)
	Delete(e *entity.Staff) (int32, error)
	Get(id int32) (*entity.Staff, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateStaff(fullname string, email string, phone string, username string, passwordHash string, createdBy int32) (int32, error)
	UpdateStaff(e *entity.Staff) (int32, error)
	DeleteStaff(e *entity.Staff) (int32, error)
	GetStaff(id int32) (*entity.Staff, error)
}
