package entity

import (
	"errors"
	"library/package/log"
	"time"
)

type Country struct {
	ID        int32
	Name      string
	IsoCode   string
	PhoneCode int16
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
}

func NewCountry(name, isoCode string, phoneCode int16, createdBy int32) (*Country, error) {
	country := &Country{
		Name:      name,
		IsoCode:   isoCode,
		PhoneCode: phoneCode,
		IsActive:  true,
		CreatedBy: createdBy,
	}
	err := country.ValidateCreate()
	if err != nil {
		log.Errorf("error validating Country entity %v", err)
		return nil, err
	}
	return country, nil
}

func (c *Country) ValidateCreate() error {
	if c.Name == "" {
		return errors.New("error validating Country entity, name field required")
	}
	if c.IsoCode == "" {
		return errors.New("error validating Country entity, iso_code field required")
	}
	if c.PhoneCode <= 0 {
		return errors.New("error validating Country entity, phone_code field required")
	}
	if c.CreatedBy <= 0 {
		return errors.New("error validating Country entity, created_by field required")
	}
	return nil
}

func (c *Country) ValidateUpdate() error {
	if c.ID <= 0 {
		return errors.New("error validating Country entity, ID field required")
	}
	if c.Name == "" {
		return errors.New("error validating Country entity, name field required")
	}
	if c.IsoCode == "" {
		return errors.New("error validating Country entity, iso_code field required")
	}
	if c.PhoneCode <= 0 {
		return errors.New("error validating Country entity, phone_code field required")
	}
	if c.UpdatedBy <= 0 {
		return errors.New("error validating Country entity, updated_by field required")
	}
	return nil
}

type CountryFilter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Name      string
}
