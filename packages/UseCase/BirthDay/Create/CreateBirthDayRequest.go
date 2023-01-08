package Create

import (
  "github.com/go-playground/validator/v10"
  "time"
  "fmt"
  "birthday-reminder/packages/Domain/Domain/Response"
  "birthday-reminder/config"
)

type CreateBirthDayRequest struct {
  UserId int  `json:"user_id"`
  Name string `json:"name"`
  Date string `json:"date"`
}

type CreateBirthDayRequestValidate struct {
  UserId int  `validate:"required"`
  Name string `validate:"required"`
  Date string `validate:"required,is_date"`
}

type CreateBirthDayRequestValidates struct {
  Elements []CreateBirthDayRequestValidate `validate:"required,dive,required"`
}


var validate *validator.Validate

func IsValidRequestBody(input_data []CreateBirthDayRequest) (error_array []Response.Error) {
  var valid_format []CreateBirthDayRequestValidate
  for _, v := range input_data {
    valid_format = append(valid_format, CreateBirthDayRequestValidate{
      UserId: v.UserId,
      Name: v.Name,
      Date: v.Date,
    })
  }
  validate = validator.New()
  validate.RegisterValidation("is_date", IsDate)
  return validateStruct(valid_format)
}

func validateStruct(valid_format []CreateBirthDayRequestValidate) (error_array []Response.Error) {
  valid_formats := CreateBirthDayRequestValidates{Elements: valid_format}
	err := validate.Struct(valid_formats)
  if err != nil {
    if _, ok := err.(*validator.InvalidValidationError); ok {
      return
    }
    cfg := config.NewConfig()
    for _, err := range err.(validator.ValidationErrors) {
      error_array = append(error_array, Response.Error{
        Message: fmt.Sprintf(cfg.VALIDATE_REQUEST_BODY_MESSAGE, err.StructNamespace(), err.Value()),
        Property: fmt.Sprintf(cfg.VALIDATE_REQUEST_BODY_PROPERTY, err.Field()),
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
