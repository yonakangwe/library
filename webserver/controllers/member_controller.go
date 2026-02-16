package controllers

import (
	"library/package/log"
	"library/package/models"
	"library/package/util"
	"library/package/validator"
	"library/package/wrappers"
	"library/services/usecase/member"

	"net/http"

	"github.com/labstack/echo/v4"
)

// const (
// 	internalServerErrorMsg = "unexpected error occurred while processing your request"
// 	recordExistsMsg        = "record already exists"
// )

// func ListRole(c echo.Context) error {

// 	m := &models.RoleFilter{}
// 	if err := c.Bind(&m); util.IsError(err) {
// 		log.Errorf("error binding pagination filter : %v", err)
// 		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
// 	}

// 	if m.Page == 0 {
// 		m.Page = 1
// 	}
// 	if m.PageSize == 0 {
// 		m.PageSize = 10
// 	}

// 	filter := &entity.RoleFilter{
// 		Page:      m.Page,
// 		PageSize:  m.PageSize,
// 		SortBy:    m.SortBy,
// 		SortOrder: m.SortOrder,
// 	}

// 	service := role.NewService()
// 	data, totalCount, err := service.List(filter)
// 	if err != nil {
// 		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
// 	}
// 	if data == nil {
// 		return wrappers.Response(c, http.StatusOK, data)
// 	}

// 	roleData := make([]*models.Role, 0)
// 	for _, d := range data {
// 		roleData = append(roleData, &models.Role{
// 			ID:          d.ID,
// 			Name:        d.Name,
// 			Description: d.Description,
// 			CreatedAt:   d.CreatedAt,
// 			UpdatedAt:   d.UpdatedAt,
// 		})
// 	}

// 	pagination := pagination.GetMetaData(filter.Page, filter.PageSize, totalCount)
// 	return wrappers.PaginationResponse(c, http.StatusOK, roleData, pagination)
// }

// func GetRole(c echo.Context) error {
// 	m := &models.Role{}
// 	if err := c.Bind(&m); util.IsError(err) {
// 		log.Errorf("error binding model id: %v", err)
// 		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
// 	}

// 	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
// 	if err := customValidator.ValidateStructPartial(m, "ID"); err != nil {
// 		log.Errorf("error validating model id: %v", err)
// 		return wrappers.ValidationErrorResponse(c, err)
// 	}

// 	service := role.NewService()
// 	data, err := service.Get(m.ID)

// 	if err != nil {
// 		log.Errorf("error getting %v: %v", m.Name, err)
// 		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
// 	}

// 	if data == nil {
// 		return wrappers.Response(c, http.StatusOK, data)
// 	}

// 	dataResponse := models.Role{
// 		ID:          data.ID,
// 		Name:        data.Name,
// 		Description: data.Description,
// 		CreatedBy:   data.CreatedBy,
// 		CreatedAt:   data.CreatedAt,
// 		UpdatedBy:   data.UpdatedBy,
// 		UpdatedAt:   data.UpdatedAt,
// 	}
// 	return wrappers.Response(c, http.StatusOK, dataResponse)
// }

func CreateMember(c echo.Context) error {
	m := &models.Member{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Role : %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	m.CreatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "FullName", "Email", "CreatedBy"); err != nil {
		log.Errorf("error validating Member  model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := member.NewService()
	_, err := service.Create(m.FullName, m.Email, m.CreatedBy)
	if util.IsError(err) {
		log.Errorf("error creating new %v: %v", m.FullName, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusCreated, " created successfully")
}
