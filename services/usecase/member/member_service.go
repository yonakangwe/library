package member

import (
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewMember()
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(fullName string, phone string, email string, membershipNo int32, createdBy int32) (int32, error) {

	member, err := entity.NewMember(fullName, phone, email, membershipNo, createdBy)
	if err != nil {
		return member.ID, err
	}
	memberID, err := s.repo.Create(member)
	if err != nil {
		return member.ID, err
	}
	return memberID, err
}
