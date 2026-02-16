package controllers

import (
	"library/package/log"
	"library/package/models"
	"library/package/pagination"
	"library/package/util"
	"library/package/validator"
	"library/package/wrappers"
	"library/services/entity"
	"library/services/usecase/country"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListCountry(c echo.Context) error {

	m := &models.CountryFilter{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding pagination filter : %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	if m.Page == 0 {
		m.Page = 1
	}
	if m.PageSize == 0 {
		m.PageSize = 10
	}

	filter := &entity.CountryFilter{
		Page:      m.Page,
		PageSize:  m.PageSize,
		SortBy:    m.SortBy,
		SortOrder: m.SortOrder,
		Name:      m.Name,
	}

	service := country.NewService()
	data, totalCount, err := service.List(filter)
	if err != nil {
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	if data == nil {
		return wrappers.Response(c, http.StatusOK, data)
	}

	countryData := make([]*models.Country, 0)
	for _, d := range data {
		countryData = append(countryData, &models.Country{
			ID:        d.ID,
			Name:      d.Name,
			IsoCode:   d.IsoCode,
			PhoneCode: d.PhoneCode,
			IsActive:  d.IsActive,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	paginationMeta := pagination.GetMetaData(filter.Page, filter.PageSize, totalCount)
	return wrappers.PaginationResponse(c, http.StatusOK, countryData, paginationMeta)
}

func GetCountry(c echo.Context) error {
	m := &models.Country{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID"); err != nil {
		log.Errorf("error validating model id: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := country.NewService()
	data, err := service.Get(m.ID)

	if err != nil {
		log.Errorf("error getting %v: %v", m.Name, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	if data == nil {
		return wrappers.Response(c, http.StatusOK, data)
	}

	dataResponse := models.Country{
		ID:        data.ID,
		Name:      data.Name,
		IsoCode:   data.IsoCode,
		PhoneCode: data.PhoneCode,
		IsActive:  data.IsActive,
		CreatedBy: data.CreatedBy,
		CreatedAt: data.CreatedAt,
		UpdatedBy: data.UpdatedBy,
		UpdatedAt: data.UpdatedAt,
	}
	return wrappers.Response(c, http.StatusOK, dataResponse)
}

func CreateCountry(c echo.Context) error {
	m := &models.Country{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Country : %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	m.CreatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "Name", "IsoCode", "PhoneCode", "CreatedBy"); err != nil {
		log.Errorf("error validating Country model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := country.NewService()
	_, err := service.Create(m.Name, m.IsoCode, m.PhoneCode, m.CreatedBy)
	if util.IsError(err) {
		log.Errorf("error creating new %v: %v", m.Name, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusCreated, m.Name+" created successfully")
}

func UpdateCountry(c echo.Context) error {
	m := models.Country{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding Country model: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	m.UpdatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID", "Name", "IsoCode", "PhoneCode", "UpdatedBy"); err != nil {
		log.Errorf("error validating Country model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}
	service := country.NewService()
	data := &entity.Country{
		ID:        m.ID,
		Name:      m.Name,
		IsoCode:   m.IsoCode,
		PhoneCode: m.PhoneCode,
		IsActive:  m.IsActive,
		UpdatedBy: m.UpdatedBy,
	}

	_, err := service.Update(data)
	if util.IsError(err) {
		log.Errorf("error updating existing %v: %v", m.Name, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, m.Name+" updated successfully")
}

func SoftDeleteCountry(c echo.Context) error {
	m := &models.Country{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding Country model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	m.DeletedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID", "DeletedBy"); err != nil {
		log.Errorf("error validating Country model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := country.NewService()
	err := service.SoftDelete(m.ID, m.DeletedBy)
	if util.IsError(err) {
		log.Errorf("error deleting record: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, "record deleted successfully")
}

func DestroyCountry(c echo.Context) error {

	m := &models.Country{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding Country model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID"); err != nil {
		log.Errorf("error validating Country model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := country.NewService()
	err := service.HardDelete(m.ID)
	if util.IsError(err) {
		log.Errorf("error deleting record: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, "record deleted successfully")
}
