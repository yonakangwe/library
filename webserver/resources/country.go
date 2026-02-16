package resources

import (
	"library/package/models"
	"library/services/entity"
)

func MapCountryEntityToModel(e *entity.Country) *models.Country {
	if e == nil {
		return nil
	}
	return &models.Country{
		ID:        e.ID,
		Name:      e.Name,
		IsoCode:   e.IsoCode,
		PhoneCode: e.PhoneCode,
		IsActive:  e.IsActive,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func MapCountryModelToEntity(m *models.Country) *entity.Country {
	if m == nil {
		return nil
	}
	return &entity.Country{
		ID:        m.ID,
		Name:      m.Name,
		IsoCode:   m.IsoCode,
		PhoneCode: m.PhoneCode,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
