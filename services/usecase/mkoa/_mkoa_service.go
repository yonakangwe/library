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
package mkoa

import (
	"library/package/log"
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewMkoa()
	return &Service{
		repo: repo,
	}
}

/*
   CREATE
*/
func (s *Service) Create(name string, code string, createdBy int32) (int32, error) {

	mkoa, err := entity.NewMkoa(name, code, int64(createdBy))
	if err != nil {
		return int32(mkoa.ID), err
	}

	mkoaID, err := s.repo.Create(mkoa)
	if err != nil {
		return int32(mkoa.ID), err
	}

	return mkoaID, err
}

/*
   LIST
*/
func (s *Service) List(filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error) {

	mkoaData, totalCount, err := s.repo.List(filter)
	if err != nil {
		return nil, 0, err
	}

	return mkoaData, totalCount, nil
}

/*
   GET
*/
func (s *Service) Get(id int32) (*entity.Mkoa, error) {

	mkoaData, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return mkoaData, nil
}

/*
   UPDATE
*/
func (s *Service) Update(e *entity.Mkoa) (int32, error) {

	err := e.ValidateUpdate()
	if err != nil {
		log.Error(err)
		return int32(e.ID), err
	}

	err = s.repo.Update(e)
	if err != nil {
		return int32(e.ID), err
	}

	return int32(e.ID), nil
}

/*
   SOFT DELETE
*/
func (s *Service) SoftDelete(id, deletedBy int32) error {

	_, err := s.Get(id)
	if err != nil {
		return err
	}

	err = s.repo.SoftDelete(id, deletedBy)
	if err != nil {
		return err
	}

	return err
}

/*
   HARD DELETE
*/
func (s *Service) HardDelete(id int32) error {

	_, err := s.Get(id)
	if err != nil {
		return err
	}

	err = s.repo.HardDelete(id)
	if err != nil {
		return err
	}

	return err
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
