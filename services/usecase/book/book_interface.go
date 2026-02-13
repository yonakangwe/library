package book

import (
	"library/services/entity"
)

type Reader interface {
}

type Writer interface {
	Create(e *entity.Book) (int32, error)
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	Create(title, author, isbn, status string, createdBy int32) (int32, error)
}
