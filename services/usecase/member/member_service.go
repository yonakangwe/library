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

func (s *Service) Create(fullName string, email string, createdBy int32) (int32, error) {
	member, err := entity.NewMember(fullName, email, createdBy)
	if err != nil {
		return member.ID, err
	}
	memberID, err := s.repo.Create(member)
	if err != nil {
		return member.ID, err
	}
	return memberID, err

}

// func (s *Service) List(filter *entity.MemberFilter) ([]*entity.Member, int32, error) {
// 	memberData, totalCount, err := s.repo.List(filter)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return memberData, totalCount, nil
// }

// func (s *Service) Get(id int32) (*entity.Member, error) {
// 	memberData, err := s.repo.Get(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return memberData, nil
// }

// func (s *Service) Update(e *entity.Member) (int32, error) {
// 	err := e.ValidateUpdate()
// 	if err != nil {
// 		log.Error(err)
// 		return e.ID, err
// 	}
// 	err = s.repo.Update(e)
// 	if err != nil {
// 		return e.ID, err
// 	}
// 	return e.ID, nil
// }

// func (s *Service) SoftDelete(id, deletedBy int32) error {
// 	_, err := s.Get(id)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.repo.SoftDelete(id, deletedBy)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// func (s *Service) HardDelete(id int32) error {
// 	_, err := s.Get(id)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.repo.HardDelete(id)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
