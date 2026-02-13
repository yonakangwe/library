package entity

import (
	"time"
	"errors"
)

type University struct {
	ID				int32
	Name			string
	Abbreviation	string
	Email			string
	Website			string
	EstablishedYear	int16
	IsActive		bool
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		time.Time
	CreatedBy		int32
	UpdatedBy		int32
	DeletedBy		int32
}

func UniversityAction(name, abbreviation, email, website, operation string, establishedYear int16, createdBy int32, updatedBy int32, deletedBy int32) (*University, error) {
	university := &University{
		Name: name,
		Abbreviation: abbreviation,
		Email: email,
		Website: website,
		EstablishedYear: establishedYear,
		IsActive: true,
	}

	if operation == "create" {
		university.CreatedAt = time.Now()
		university.CreatedBy = createdBy
	}

	if operation == "update" {
		university.UpdatedBy = updatedBy
		university.UpdatedAt = time.Now()
	}

	if operation == "delete" {
		university.DeletedBy = deletedBy
		university.DeletedAt = time.Now()
	}

	// Generate abbreaviation if not provided
	if university.Abbreviation == "" {
		university.Abbreviation = university.generateAbbreviation()
	}

	err := university.ValidateFields(operation)
	if err != nil {
		return nil, err
	}

	// Return the university entity
	return university, nil;
}

func (university *University) ValidateFields(operation string) error {
	if university.Name == "" {
		return errors.New("Name is required")
	}

	if university.Abbreviation == "" {
		return errors.New("Abbreviation is required")
	}

	if university.Email == "" {
		return errors.New("Email is required")
	}

	if university.Website == "" {
		return errors.New("Website is required")
	}

	if university.EstablishedYear <= 0 {
		return errors.New("Established year is required")
	}

	if operation == "create" && university.CreatedBy <= 0 {
		return errors.New("Created by is required")
	}

	if operation == "update" && university.UpdatedBy <= 0 {
		return errors.New("Updated by is required")
	}

	if operation == "delete" && university.DeletedBy <= 0 {
		return errors.New("Deleted by is required")
	}

	return nil;
}

func (university *University) generateAbbreviation() string {
	if university.Name == "" {
		return ""
	}
	return university.Name[:4]
}

type UniversityFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Name      string
}