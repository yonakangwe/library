package staff

import (
	"library/services/entity"
	"library/services/repository"
	// "github.com/k0kubun/pp"
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
	//pp.Println(staff)
	if err != nil {
		return staff.ID, err
	}
	staffID, err := s.repo.Create(staff)
	if err != nil {
		return staff.ID, err
	}
	return staffID, err

}
