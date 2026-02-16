package controllers

import (
	"library/package/log"
	"library/package/models"
	"library/package/util"
	"library/package/validator"
	"library/package/wrappers"
	"library/services/entity"
	"library/services/usecase/staff"
	"strconv"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateStaff(c echo.Context) error {
	m := &models.Staff{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Role : %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	m.CreatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "FullName", "Email", "Phone", "Username", "PasswordHash", "CreatedBy"); err != nil {
		log.Errorf("error validating Role  model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := staff.NewService()
	_, err := service.CreateStaff(m.FullName, m.Email, m.Phone, m.Username, m.PasswordHash, m.CreatedBy)
	if util.IsError(err) {
		log.Errorf("error creating new %v: %v", m.FullName, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusCreated, "Data saved successfully")
}

func UpdateStaff(c echo.Context) error {
	// HATUA 1: Pata ID kutoka URL parameter
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return wrappers.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	// HATUA 2: Bind JSON request body
	m := models.Staff{}
	if err := c.Bind(&m); util.IsError(err) {
		log.Errorf("error binding staff model: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	// HATUA 3: Weka ID kutoka URL na updated by
	m.ID = int32(id)
	m.UpdatedBy = 1

	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "FullName", "Username", "UpdatedBy"); err != nil {
		log.Errorf("error validating Staff  model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := staff.NewService()
	updatedBy := m.UpdatedBy // Convert to pointer
	data := &entity.Staff{
		ID:           m.ID,
		FullName:     m.FullName,
		Email:        m.Email,
		Phone:        m.Phone,
		Username:     m.Username,
		PasswordHash: m.PasswordHash,
		UpdatedBy:    &updatedBy, // 👈 Pointer
	}

	_, err = service.UpdateStaff(data)
	if util.IsError(err) {
		log.Errorf("error updating existing %v: %v", m.FullName, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, m.FullName+" updated successfully")
}

func DestroyStaff(c echo.Context) error {
	// HATUA 1: Pata ID kutoka URL parameter
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return wrappers.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	service := staff.NewService()
	_, err = service.DeleteStaff(&entity.Staff{ID: int32(id)})
	if util.IsError(err) {
		log.Errorf("error deleting record: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, "record deleted successfully")
}

func GetStaff(c echo.Context) error {
	// HATUA 1: Pata ID kutoka URL parameter
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return wrappers.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	service := staff.NewService()
	data, err := service.GetStaff(int32(id))

	if err != nil {
		log.Errorf("error getting staff by id %v: %v", id, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
	}

	if data == nil {
		return wrappers.Response(c, http.StatusNotFound, "Staff not found")
	}

	// HATUA 2: Usirudishe passwordHash kwa usalama
	// Tengeneza default values kwa pointers
	updatedBy := int32(0)
	if data.UpdatedBy != nil {
		updatedBy = *data.UpdatedBy
	}

	updatedAt := time.Time{}
	if data.UpdatedAt != nil {
		updatedAt = *data.UpdatedAt
	}

	dataResponse := models.Staff{
		ID:           data.ID,
		FullName:     data.FullName,
		Email:        data.Email,
		Phone:        data.Phone,
		Username:     data.Username,
		PasswordHash: "", // 👈 Usirudishe password
		CreatedBy:    data.CreatedBy,
		CreatedAt:    data.CreatedAt,
		UpdatedBy:    updatedBy, // 👈 Sasa ni int32
		UpdatedAt:    updatedAt, // 👈 Sasa ni time.Time
	}
	return wrappers.Response(c, http.StatusOK, dataResponse)
}
