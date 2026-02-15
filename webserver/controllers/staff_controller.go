package controllers

import (
	"library/package/log"
	"library/package/models"
	"library/package/util"
	"library/package/validator"
	"library/package/wrappers"
	"library/services/usecase/staff"

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
	return wrappers.MessageResponse(c, http.StatusCreated, m.FullName+" created successfully")
}
