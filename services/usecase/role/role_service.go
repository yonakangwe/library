package role

import (
	"library/package/log"
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewRole()
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name string, description string, createdBy int32) (int32, error) {
	role, err := entity.NewRole(name, description, createdBy)
	if err != nil {
		return role.ID, err
	}
	roleID, err := s.repo.Create(role)
	if err != nil {
		return role.ID, err
	}
	return roleID, err

}

func (s *Service) List(filter *entity.RoleFilter) ([]*entity.Role, int32, error) {
	roleData, totalCount, err := s.repo.List(filter)
	if err != nil {
		return nil, 0, err
	}
	return roleData, totalCount, nil
}

func (s *Service) Get(id int32) (*entity.Role, error) {
	roleData, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return roleData, nil
}

func (s *Service) Update(e *entity.Role) (int32, error) {
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
