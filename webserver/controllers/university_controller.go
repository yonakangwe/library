package controllers

import (
	"library/package/log"
	"library/package/models"
	"library/package/wrappers"
	"library/services/entity"
	"library/services/usecase/university"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateUniversity(ctx echo.Context) error {
	var model models.University

	// Bind request body to model
	if err := ctx.Bind(&model); err != nil {
		log.Error("Invalid request body: %v", err)
		return wrappers.Response(ctx, http.StatusBadRequest, "Invalid request format")
	}

	// Validate model
	validate := validator.New()
	if err := validate.Struct(model); err != nil {
		log.Error("Validation error: %v", err)
		return wrappers.Response(ctx, http.StatusBadRequest, err.Error())
	}

	// Map model to entity
	universityEntity := &entity.University{
		Name:            model.Name,
		Abbreviation:    model.Abbreviation,
		Email:           model.Email,
		Website:         model.Website,
		EstablishedYear: model.EstablishedYear,
		IsActive:        model.IsActive,
		CreatedBy:       model.CreatedBy,
		UpdatedBy:       model.UpdatedBy,
		DeletedBy:       model.DeletedBy,
	}

	service := university.UniversityService()
	ID, err := service.CreateUniversity(universityEntity)

	if err != nil {
		log.Error("Error creating university: %v", err)
		return wrappers.Response(ctx, http.StatusInternalServerError, "Failed to create university")
	}

	log.Info("University created successfully with ID: %d", ID)
	return wrappers.Response(ctx, http.StatusCreated, "University created successfully")
}
