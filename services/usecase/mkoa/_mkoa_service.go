package mkoa

import (
	"context"
	"errors"

	"library/services/entity"
)

type service struct {
	repo Repository
}

func NewUsecase(repo Repository) Usecase {
	return &service{repo: repo}
}

/*
CREATE
*/
func (s *service) Create(ctx context.Context, name, code string, createdBy int64) (*entity.Mkoa, error) {

	m, err := entity.NewMkoa(name, code, createdBy)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Insert(ctx, m); err != nil {
		return nil, err
	}

	return m, nil
}

/*
GET ALL
*/
func (s *service) GetAll(ctx context.Context) ([]*entity.Mkoa, error) {
	return s.repo.GetAll(ctx)
}

/*
GET BY ID
*/
func (s *service) GetByID(ctx context.Context, id int64) (*entity.Mkoa, error) {
	return s.repo.GetByID(ctx, id)
}

/*
UPDATE
*/
func (s *service) Update(ctx context.Context, m *entity.Mkoa) error {
	return s.repo.Update(ctx, m)
}

/*
DELETE (Soft Delete)
*/
func (s *service) Delete(ctx context.Context, id int64, deletedBy int64) error {

	if id <= 0 {
		return errors.New("invalid id")
	}
	if deletedBy <= 0 {
		return errors.New("deleted_by required")
	}

	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	m.DeletedBy = deletedBy

	return s.repo.SoftDelete(ctx, m)
}
