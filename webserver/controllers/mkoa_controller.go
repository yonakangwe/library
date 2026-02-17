package controllers

import (
	"errors"
	"library/package/log"
	"library/package/models"
	"library/package/pagination"
	"library/package/report"
	"library/package/util"
	"library/package/validator"
	"library/package/wrappers"
	"library/services/entity"
	"library/services/usecase/mkoa"
	"strconv"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	mkoaInternalServerErrorMsg = "unexpected error occurred while processing your request"
	mkoaDBUnavailableMsg       = "Database unavailable. Please try again later."
)

func CreateMkoa(c echo.Context) error {
	m := &models.Mkoa{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	m.CreatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "Name", "Code", "CreatedBy"); err != nil {
		log.Errorf("error validating Mkoa model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := mkoa.NewService()
	mkoaID, err := service.Create(m.Name, m.Code, m.CreatedBy)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		if errors.Is(err, mkoa.ErrCodeExists) {
			return wrappers.ErrorResponse(c, http.StatusConflict, "Code already exists. Please use a unique code.")
		}
		log.Errorf("error creating new mkoa %v: %v", m.Name, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}
	return wrappers.Response(c, http.StatusCreated, map[string]any{
		"id":      mkoaID,
		"message": m.Name + " created successfully",
	})
}

func ListMkoa(c echo.Context) error {
	m := &models.MkoaFilter{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa filter: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	if m.Page == 0 {
		m.Page = 1
	}
	if m.PageSize == 0 {
		m.PageSize = 10
	}

	filter := &entity.MkoaFilter{
		Page:      m.Page,
		PageSize:  m.PageSize,
		SortBy:    m.SortBy,
		SortOrder: m.SortOrder,
		Name:      m.Name,
		Code:      m.Code,
		Status:    m.Status,
	}

	service := mkoa.NewService()
	data, totalCount, err := service.List(filter)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		log.Errorf("error listing mkoa: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}
	if data == nil {
		return wrappers.Response(c, http.StatusOK, []*models.Mkoa{})
	}

	mkoaData := make([]*models.Mkoa, 0, len(data))
	for _, d := range data {
		updatedAt := time.Time{}
		if d.UpdatedAt != nil {
			updatedAt = *d.UpdatedAt
		}
		deletedAt := time.Time{}
		if d.DeletedAt != nil {
			deletedAt = *d.DeletedAt
		}
		mkoaData = append(mkoaData, &models.Mkoa{
			ID:        int32(d.ID),
			Name:      d.Name,
			Code:      d.Code,
			Status:    d.Status,
			CreatedBy: int32(entity.Int64PtrVal(d.CreatedBy)),
			UpdatedBy: int32(entity.Int64PtrVal(d.UpdatedBy)),
			DeletedBy: int32(entity.Int64PtrVal(d.DeletedBy)),
			CreatedAt: d.CreatedAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}

	meta := pagination.GetMetaData(filter.Page, filter.PageSize, totalCount)
	return wrappers.PaginationResponse(c, http.StatusOK, mkoaData, meta)
}

func GetMkoa(c echo.Context) error {
	m := &models.Mkoa{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID"); err != nil {
		log.Errorf("error validating Mkoa model id: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := mkoa.NewService()
	data, err := service.Get(m.ID)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrNotFound) {
			return wrappers.ErrorResponse(c, http.StatusNotFound, "mkoa not found")
		}
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		log.Errorf("error getting mkoa %v: %v", m.ID, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	if data == nil {
		return wrappers.ErrorResponse(c, http.StatusNotFound, "mkoa not found")
	}

	updatedAt := time.Time{}
	if data.UpdatedAt != nil {
		updatedAt = *data.UpdatedAt
	}
	deletedAt := time.Time{}
	if data.DeletedAt != nil {
		deletedAt = *data.DeletedAt
	}
	dataResponse := models.Mkoa{
		ID:        int32(data.ID),
		Name:      data.Name,
		Code:      data.Code,
		Status:    data.Status,
		CreatedBy: int32(entity.Int64PtrVal(data.CreatedBy)),
		UpdatedBy: int32(entity.Int64PtrVal(data.UpdatedBy)),
		DeletedBy: int32(entity.Int64PtrVal(data.DeletedBy)),
		CreatedAt: data.CreatedAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
	return wrappers.Response(c, http.StatusOK, dataResponse)
}

func UpdateMkoa(c echo.Context) error {
	m := &models.Mkoa{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa model: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	m.UpdatedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID", "Name", "Code", "UpdatedBy"); err != nil {
		log.Errorf("error validating Mkoa model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	status := m.Status
	if status == "" {
		status = entity.MkoaStatusActive
	}
	ub := int64(m.UpdatedBy)
	e := &entity.Mkoa{
		ID:        int64(m.ID),
		Name:      m.Name,
		Code:      m.Code,
		Status:    status,
		UpdatedBy: &ub,
	}

	service := mkoa.NewService()
	_, err := service.Update(e)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrNotFound) {
			return wrappers.ErrorResponse(c, http.StatusNotFound, "mkoa not found")
		}
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		if errors.Is(err, mkoa.ErrCodeExists) {
			return wrappers.ErrorResponse(c, http.StatusConflict, "Code already exists. Please use a unique code.")
		}
		log.Errorf("error updating mkoa %v: %v", m.Name, err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, m.Name+" updated successfully")
}

func SoftDeleteMkoa(c echo.Context) error {
	m := &models.Mkoa{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	m.DeletedBy = 1
	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID", "DeletedBy"); err != nil {
		log.Errorf("error validating Mkoa model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := mkoa.NewService()
	err := service.SoftDelete(m.ID, m.DeletedBy)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrNotFound) {
			return wrappers.ErrorResponse(c, http.StatusNotFound, "mkoa not found")
		}
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		log.Errorf("error deleting mkoa record: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, "record deleted successfully")
}

func DestroyMkoa(c echo.Context) error {
	m := &models.Mkoa{}
	if err := c.Bind(m); util.IsError(err) {
		log.Errorf("error binding Mkoa model id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
	if err := customValidator.ValidateStructPartial(m, "ID"); err != nil {
		log.Errorf("error validating Mkoa model: %v", err)
		return wrappers.ValidationErrorResponse(c, err)
	}

	service := mkoa.NewService()
	err := service.HardDelete(m.ID)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrNotFound) {
			return wrappers.ErrorResponse(c, http.StatusNotFound, "mkoa not found")
		}
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		log.Errorf("error destroying mkoa record: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}
	return wrappers.MessageResponse(c, http.StatusAccepted, "record deleted successfully")
}

// ExportMkoaReport generates a simple PDF report of all mkoa records and returns it as a download.
func ExportMkoaReport(c echo.Context) error {
	// Use a large page size to fetch all items (only non-deleted).
	filter := &entity.MkoaFilter{
		Page:      1,
		PageSize:  1000,
		SortBy:    "name",
		SortOrder: "ASC",
	}

	service := mkoa.NewService()
	data, _, err := service.List(filter)
	if util.IsError(err) {
		if errors.Is(err, mkoa.ErrDBUnavailable) {
			return wrappers.ErrorResponse(c, http.StatusServiceUnavailable, mkoaDBUnavailableMsg)
		}
		log.Errorf("error listing mkoa for pdf report: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	// Build table data: first row is header.
	reportData := make([][]string, 0, len(data)+1)
	reportData = append(reportData, []string{"#", "Name", "Code", "Status", "Created at"})

	for i, d := range data {
		createdAt := d.CreatedAt.Format("2006-01-02 15:04")
		reportData = append(reportData, []string{
			strconv.Itoa(i + 1),
			d.Name,
			d.Code,
			d.Status,
			createdAt,
		})
	}

	// Reasonable column widths (percentages, will be normalised by report package).
	columnWidth := []float64{8, 34, 18, 15, 25}

	path := report.GeneralReport(
		"SQA System",
		"Mikoa (Regions)",
		reportData,
		columnWidth,
		"mikoa_regions",
		9,
		false,
	)
	if path == "" {
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, mkoaInternalServerErrorMsg)
	}

	return c.Attachment(path, "mikoa_regions.pdf")
}
