package Create

import (
  "github.com/go-playground/validator/v10"
  "time"
  "fmt"
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Response"
  "birthday-reminder/config"
  domain_auth "birthday-reminder/packages/Domain/Domain/Auth"
  domain_user "birthday-reminder/packages/Domain/Domain/User"
  "birthday-reminder/packages/Domain/Domain/Request"
)

var validate *validator.Validate

type CreateBirthDayRequest struct {
  UserId int64  `json:"user_id"`
  Name string `json:"name"`
  Date string `json:"date"`
}

type CreateBirthDayRequestValidate struct {
  UserId int64  `validate:"required"`
  Name string `validate:"required"`
  Date string `validate:"required,is_date"`
}

type CreateBirthDayRequestValidates struct {
  Elements []CreateBirthDayRequestValidate `validate:"required,dive,required"`
}

func GetRequestBody(w http.ResponseWriter, r *http.Request) (input_data2 []CreateBirthDayRequest, errors []Response.Error) {
  // json decode
  input_data := []CreateBirthDayRequest{}
  Request.JsonDecode(r, &input_data)
  user_id := domain_auth.GetUserIdByToken(w, r)
  input_data2 = []CreateBirthDayRequest{}
  input_data2 = convertCreateInputData(input_data2, input_data, user_id)
  if errors := validRequestBody(input_data2); len(errors) > 0 {
    return nil, errors
  }
  return input_data2, nil
}
// Let the user_id bind to the object
func convertCreateInputData(
  result []CreateBirthDayRequest,
  target []CreateBirthDayRequest,
  user_id domain_user.UserId,
) ([]CreateBirthDayRequest) {
  for _, v := range target {
    result = append(result, CreateBirthDayRequest{
      UserId: user_id.GetValue(),
      Name: v.Name,
      Date: v.Date,
    })
  }
  return result
}

// isValidRequestBody() check request body
func validRequestBody(input_data []CreateBirthDayRequest) (error_array []Response.Error) {
  var valid_format []CreateBirthDayRequestValidate
  for _, v := range input_data {
    valid_format = append(valid_format, CreateBirthDayRequestValidate{
      UserId: v.UserId,
      Name: v.Name,
      Date: v.Date,
    })
  }
  validate = validator.New()
  validate.RegisterValidation("is_date", isDate)
  return validateStruct(valid_format)
}
// isDate() is valid Date Format for create birth day?
func isDate(fl validator.FieldLevel) bool {
  _, err := time.Parse("2006-01-02", fl.Field().String())
  if err != nil {
    return false
  }
  return true
}

// validateStruct() is valid struct?
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


