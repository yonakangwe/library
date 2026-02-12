package resources

import (
	"library/package/models"
	"library/services/entity"
)

func MapRoleEntityToModel(e *entity.Role) *models.Role {
	if e == nil {
		return nil
	}
	return &models.Role{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		DeletedBy:   e.DeletedBy,
		DeletedAt:   e.DeletedAt,
	}
}

func MapRoleModelToEntity(m *models.Role) *entity.Role {
	if m == nil {
		return nil
	}

	return &entity.Role{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		DeletedBy:   m.DeletedBy,
		DeletedAt:   m.DeletedAt,
	}
}
