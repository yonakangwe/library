package validator

import "github.com/go-playground/validator/v10"

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// JSONErrors struct
type JSONErrors struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

// Validate Helper
func Validate(s interface{}) []JSONErrors {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		errors := []JSONErrors{}
		if _, ok := err.(*validator.InvalidValidationError); ok {
			println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			n := JSONErrors{Field: err.Field(), Rule: err.ActualTag()}
			errors = append(errors, n)
		}
		return errors
	}
	return nil
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (cv *CustomValidator) ValidateStructPartial(i interface{}, fields ...string) error {
	return cv.validator.StructPartial(i, fields...)
}

// GetValidator returns a custom validator
func GetValidator() *CustomValidator {
	cv := &CustomValidator{validator: validator.New()}
	return cv
}
