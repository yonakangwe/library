package university

import (
	"library/services/entity"
	"library/services/repository"
)

type Service struct {
	mInterface UniversityInterface
}

func UniversityService() UseCase {
	repository := repository.NewInstance()
	return &Service{
		mInterface: repository,
	}
}

func (service *Service) CreateUniversity(university *entity.University) (int32, error) {
	// Validate first through entity
	mUniversity, err := entity.UniversityAction(university, "create")
	if err != nil {
		return 0, err
	}

	ID, err := service.mInterface.Create(mUniversity)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
