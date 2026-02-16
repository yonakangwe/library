package models

type University struct {
	ID				int32 			`json:"id" validate:"numeric,required"`
	Name			string			`json:"name" validate:"required"`
	Abbreviation	string			`json:"abbreviation" validate:"required"`
	Email			string			`json:"email" validate:"required,email"`
	Website			string			`json:"website" validate:"required,url"`
	EstablishedYear	int16			`json:"established_year" validate:"numeric,required"`
	IsActive		bool			`json:"is_active" validate:"required,boolean"`
	CreatedBy		int32			`json:"created_by" validate:"numeric,required"`
	UpdatedBy		int32			`json:"updated_by" validate:"numeric,required"`
	DeletedBy		int32			`json:"deleted_by" validate:"numeric,required"`
}