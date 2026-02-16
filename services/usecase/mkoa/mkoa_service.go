package mkoa

import (
	"errors"
	"library/package/log"
	"library/services/entity"
	"library/services/repository"
)

// ErrNotFound is returned when a mkoa record is not found (e.g. Get by id).
var ErrNotFound = errors.New("mkoa not found")

// ErrDBUnavailable is returned when the database is not connected (e.g. PostgreSQL not running).
var ErrDBUnavailable = errors.New("database unavailable")

// ErrCodeExists is returned when creating or updating with a code that already exists.
var ErrCodeExists = errors.New("code already exists")

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
		if errors.Is(err, repository.ErrDBUnavailable) {
			return 0, ErrDBUnavailable
		}
		if errors.Is(err, repository.ErrMkoaCodeExists) {
			return 0, ErrCodeExists
		}
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
		if errors.Is(err, repository.ErrDBUnavailable) {
			return nil, 0, ErrDBUnavailable
		}
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
		if errors.Is(err, repository.ErrMkoaNotFound) {
			return nil, ErrNotFound
		}
		if errors.Is(err, repository.ErrDBUnavailable) {
			return nil, ErrDBUnavailable
		}
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
		if errors.Is(err, repository.ErrMkoaNotFound) {
			return 0, ErrNotFound
		}
		if errors.Is(err, repository.ErrDBUnavailable) {
			return 0, ErrDBUnavailable
		}
		if errors.Is(err, repository.ErrMkoaCodeExists) {
			return 0, ErrCodeExists
		}
		return int32(e.ID), err
	}
	return int32(e.ID), nil
}

/*
SOFT DELETE
*/
func (s *Service) SoftDelete(id, deletedBy int32) error {
	err := s.repo.SoftDelete(id, deletedBy)
	if err != nil {
		if errors.Is(err, repository.ErrMkoaNotFound) {
			return ErrNotFound
		}
		if errors.Is(err, repository.ErrDBUnavailable) {
			return ErrDBUnavailable
		}
		return err
	}
	return nil
}

/*
HARD DELETE
*/
func (s *Service) HardDelete(id int32) error {
	err := s.repo.HardDelete(id)
	if err != nil {
		if errors.Is(err, repository.ErrMkoaNotFound) {
			return ErrNotFound
		}
		if errors.Is(err, repository.ErrDBUnavailable) {
			return ErrDBUnavailable
		}
		return err
	}
	return nil
}
