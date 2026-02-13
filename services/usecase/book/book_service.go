package book

import (
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	repo Repository
}

func NewService() Usecase {
	repo := repository.NewBook()
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(title, author, isbn, status string, createdBy int32) (int32, error) {
	book, err := entity.NewBook(title, author, isbn, status, createdBy)
	if err != nil {
		return book.ID, err
	}
	bookID, err := s.repo.Create(book)
	if err != nil {
		return book.ID, err
	}
	return bookID, err
}
