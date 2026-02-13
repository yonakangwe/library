package country

import (
	"library/package/log"
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewCountry()
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name, isoCode string, phoneCode int16, createdBy int32) (int32, error) {
	country, err := entity.NewCountry(name, isoCode, phoneCode, createdBy)
	if err != nil {
		return country.ID, err
	}
	countryID, err := s.repo.Create(country)
	if err != nil {
		return country.ID, err
	}
	return countryID, err
}

func (s *Service) List(filter *entity.CountryFilter) ([]*entity.Country, int32, error) {
	countryData, totalCount, err := s.repo.List(filter)
	if err != nil {
		return nil, 0, err
	}
	return countryData, totalCount, nil
}

func (s *Service) Get(id int32) (*entity.Country, error) {
	countryData, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return countryData, nil
}

func (s *Service) Update(e *entity.Country) (int32, error) {
	err := e.ValidateUpdate()
	if err != nil {
		log.Error(err)
		return e.ID, err
	}
	err = s.repo.Update(e)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

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
