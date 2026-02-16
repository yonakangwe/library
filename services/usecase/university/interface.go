package university

import "library/services/entity"

type Writer /*Setter*/ interface {
	Create(entity *entity.University) (int32, error)
}

type UniversityInterface interface {
	Writer
}

type UseCase interface {
	CreateUniversity (entity *entity.University) (int32, error)
}