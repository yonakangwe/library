package entity

import (
	"errors"
	"library/package/log"
	"time"
)

type Book struct {
	ID        int32
	Title     string
	Author    string
	Isbn      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
}

func NewBook(title, author, isbn, status string, createdBy int32) (*Book, error) {
	book := &Book{
		Title:     title,     // Initialization of Title field
		Author:    author,    // Initialization of Author field
		Isbn:      isbn,      // Initialization of Isbn field
		Status:    status,    // Initialization of Status field
		CreatedBy: createdBy, // Initialization of CreatedBy field
	}
	err := book.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Book entity %v", err)
		return nil, err
	}
	return book, nil
}

func (b *Book) ValidateCreate() error {
	if b.Title == "" {
		return errors.New("error validating Book entity, title field required")
	}
	if b.Author == "" {
		return errors.New("error validating Book entity, author field required")
	}
	if b.Isbn == "" {
		return errors.New("error validating Book entity, isbn field required")
	}
	if b.Status == "" {
		return errors.New("error validating Book entity, status field required")
	}
	if b.CreatedBy <= 0 {
		return errors.New("error validating Book entity, created_by field required")
	}
	return nil
}

func (b *Book) ValidateUpdate() error {
	if b.ID <= 0 {
		return errors.New("error validating Book entity, ID field required")
	}
	if b.Title == "" {
		return errors.New("error validating Book entity, title field required")
	}
	if b.Author == "" {
		return errors.New("error validating Book entity, author field required")
	}
	if b.Isbn == "" {
		return errors.New("error validating Book entity, isbn field required")
	}
	if b.Status == "" {
		return errors.New("error validating Book entity, status field required")
	}
	if b.UpdatedBy <= 0 {
		return errors.New("error validating Book entity, updated_by field required")
	}
	return nil
}
