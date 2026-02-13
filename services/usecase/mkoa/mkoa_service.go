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
	return NewUsecase(repo)
}

func NewUsecase(repo Repository) UseCase {
	return &Service{repo: repo}
}

/*
CREATE
*/
func (s *Service) Create(name string, code string, createdBy int32) (int32, error) {
	mkoa, err := entity.NewMkoa(name, code, int64(createdBy))
	if err != nil {
		return 0, err
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
	return s.repo.SoftDelete(id, deletedBy)
}

/*
HARD DELETE
*/
func (s *Service) HardDelete(id int32) error {
	return s.repo.HardDelete(id)
}
