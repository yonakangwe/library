package mkoa

import (
	"context"

	"library/package/log"
	"library/services/database"
	"library/services/entity"
)

// mkoaRepoAdapter adapts database.MkoaRepository to the usecase Repository interface
type mkoaRepoAdapter struct {
	*database.MkoaRepository
}

func (a *mkoaRepoAdapter) Create(e *entity.Mkoa) (int32, error) {
	err := a.Insert(context.Background(), e)
	return int32(e.ID), err
}

func (a *mkoaRepoAdapter) Get(id int32) (*entity.Mkoa, error) {
	return a.GetByID(context.Background(), int64(id))
}

func (a *mkoaRepoAdapter) List(filter *entity.MkoaFilter) ([]*entity.Mkoa, int32, error) {
	return a.MkoaRepository.List(context.Background(), filter)
}

func (a *mkoaRepoAdapter) Update(e *entity.Mkoa) error {
	return a.MkoaRepository.Update(context.Background(), e)
}

func (a *mkoaRepoAdapter) SoftDelete(id, deletedBy int32) error {
	m, err := a.GetByID(context.Background(), int64(id))
	if err != nil {
		return err
	}
	m.DeletedBy = int64(deletedBy)
	return a.MkoaRepository.SoftDelete(context.Background(), m)
}

func (a *mkoaRepoAdapter) HardDelete(id int32) error {
	return a.MkoaRepository.HardDelete(context.Background(), int64(id))
}

type Service struct {
	repo Repository
}

func NewService() UseCase {
	db, err := database.Connect()
	if err != nil {
		log.Errorf("mkoa service: db connect failed: %v", err)
		panic(err)
	}
	repo := database.NewMkoaRepository(db)
	return NewUsecase(&mkoaRepoAdapter{MkoaRepository: repo})
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
