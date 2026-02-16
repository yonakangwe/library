package staff

import (
	"library/services/entity"
	"library/services/repository"

	"github.com/k0kubun/pp"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewStaff()
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateStaff(fullname string, email string, phone string, username string, passwordHash string, createdBy int32) (int32, error) {

	staff, err := entity.NewStaff(fullname, email, phone, username, passwordHash, createdBy)
	pp.Println(staff)
	if err != nil {
		return 0, err // 👈 FIXED: Rudisha 0 sio staff.ID
	}
	staffID, err := s.repo.Create(staff)
	if err != nil {
		return 0, err // 👈 FIXED: Rudisha 0 sio staff.ID
	}
	return staffID, err

}

func (s *Service) UpdateStaff(e *entity.Staff) (int32, error) {
	err := e.ValidateUpdate()
	if err != nil {
		return e.ID, err
	}
	return s.repo.Update(e)
}

func (s *Service) DeleteStaff(e *entity.Staff) (int32, error) {
	return s.repo.Delete(e)
}

func (s *Service) GetStaff(id int32) (*entity.Staff, error) {
	return s.repo.Get(id)
}
