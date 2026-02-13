package staff

import (
	"library/services/entity"
)

type Reader interface {
}

type Writer interface {
	Create(e *entity.Staff) (int32, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateStaff(fullname string, email string, phone string, username string, passwordHash string, createdBy int32) (int32, error)
}
