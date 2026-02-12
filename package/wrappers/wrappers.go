package wrappers

import (
	"library/package/models"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Code    int          `json:"code"`
	Error   string       `json:"error,omitempty"`
	Message string       `json:"message,omitempty"`
	Meta    *models.Meta `json:"meta,omitempty"`
	Data    any          `json:"data,omitempty"`
}

func Response(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, ResponseData{
		Code: statusCode,
		Data: data,
	})
}

func PaginationResponse(c echo.Context, statusCode int, data any, meta *models.Meta) error {
	return c.JSON(statusCode, ResponseData{
		Code:    statusCode,
		Message: "Success",
		Meta:    meta,
		Data:    data,
	})
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ResponseData{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, err string, payload ...any) error {
	var data any
	if len(payload) > 0 {
		data = payload[0]
	}
	return c.JSON(statusCode, ResponseData{
		Code:  statusCode,
		Error: err,
		Data:  data,
	})
}
func ValidationErrorResponse(c echo.Context, verr error) error {
	m := map[string]string{}

	if ve, ok := verr.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())
			switch fe.Tag() {
			case "required":
				m[field] = field + " is required"
			case "min":
				m[field] = field + " is too short"
			case "max":
				m[field] = field + " is too long"
			case "email":
				m[field] = field + " must be a valid email"
			case "len":
				m[field] = field + " length is invalid"
			case "numeric":
				m[field] = field + " must be numeric"
			case "uuid":
				m[field] = field + " must be a valid UUID"
			case "oneof":
				m[field] = field + " has an invalid value"
			case "url":
				m[field] = field + " must be a valid URL"
			case "gte":
				m[field] = field + " must be greater than or equal to the minimum"
			case "lte":
				m[field] = field + " must be less than or equal to the maximum"
			default:
				m[field] = "invalid " + field
			}

		}
	} else {
		m["form"] = "invalid input"
	}

	return ErrorResponse(c, http.StatusUnprocessableEntity, "validation error, some required fields are missing", m)
}
