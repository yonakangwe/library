package university

import "library/services/entity"

type Service struct {
	mInterface UniversityInterface
}

func (service *Service) CreateUniversity(university *entity.University) (int32, error) {
	//validate first through entity
	err := university.ValidateFields("create")
	if err != nil {
		return 0, err
	}
	
	ID, err := service.mInterface.Create(university)
	if err != nil {
		return university.ID, err
	}
	return ID, nil
}

