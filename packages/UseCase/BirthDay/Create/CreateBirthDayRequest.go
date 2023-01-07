package Create

import (
  "github.com/go-playground/validator/v10"
  "time"
  "fmt"
  "birthday-reminder/packages/Domain/Domain/Response"
)

type CreateBirthDayRequest struct {
  UserId int  `json:"user_id"`
  Date string `json:"date"`
}

type CreateBirthDayRequestValidate struct {
  UserId int  `validate:"required"`
  Date string `validate:"required,is_date"`
}

type CreateBirthDayRequestValidates struct {
  Elements []CreateBirthDayRequestValidate `validate:"required,dive,required"`
}


var validate *validator.Validate

func IsValid(req []CreateBirthDayRequestValidate) (error_array []Response.Error) {
	validate = validator.New()
  validate.RegisterValidation("is_date", IsDate)
	return validateStruct(req)
}

func validateStruct(req []CreateBirthDayRequestValidate) (error_array []Response.Error) {
  requests := CreateBirthDayRequestValidates{Elements: req}
	err := validate.Struct(requests)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return
		}
		for _, err := range err.(validator.ValidationErrors) {
      error_array = append(error_array, Response.Error{
        Message: fmt.Sprintf("Validation Error. Property:%v Value:%v", err.StructField(), err.Value()),
        Property: fmt.Sprintf("%v", err.Field()),
      })
		}
		return
	}
  return error_array
}

func IsDate(fl validator.FieldLevel) bool {
  _, err := time.Parse("2006-01-02", fl.Field().String())
  if err != nil {
    return false
  }
  return true
}
