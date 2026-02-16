package entity

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type University struct {
	ID              int32
	Name            string
	Abbreviation    string
	Email           string
	Website         string
	EstablishedYear int16
	IsActive        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	CreatedBy       int32
	UpdatedBy       int32
	DeletedBy       int32
}

func UniversityAction(university *University, operation string) (*University, error) {
	university = &University{
		ID:              university.ID,
		Name:            university.Name,
		Abbreviation:    university.Abbreviation,
		Email:           university.Email,
		Website:         university.Website,
		EstablishedYear: university.EstablishedYear,
		IsActive:        university.IsActive,
		CreatedBy:       university.CreatedBy,
		UpdatedBy:       university.UpdatedBy,
		DeletedBy:       university.DeletedBy,
	}

	if operation == "create" {
		university.CreatedAt = time.Now()
	}

	if operation == "update" {
		university.UpdatedAt = time.Now()
	}

	if operation == "delete" {
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
	return university, nil
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

	// Validate email format
	if _, err := mail.ParseAddress(university.Email); err != nil {
		return errors.New("Invalid email format")
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

	return nil
}

func (university *University) generateAbbreviation() string {
	if university.Name == "" {
		return ""
	}
	words := strings.Fields(university.Name)
	abbr := ""
	for _, word := range words {
		if len(word) > 0 {
			abbr += string(word[0])
		}
	}
	return strings.ToUpper(abbr)
}

type UniversityFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Name      string
}
