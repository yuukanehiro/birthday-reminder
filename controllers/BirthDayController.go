package controllers

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  domain_auth "birthday-reminder/packages/Domain/Domain/Auth"
  domain_user "birthday-reminder/packages/Domain/Domain/User"
  "birthday-reminder/packages/Domain/Domain/Response"
  "birthday-reminder/packages/Domain/Domain/Request"
  "net/http"
)

type BirthDayControllerInterface interface {
  ListBirthDay(w http.ResponseWriter, r *http.Request) Response.ApiResponseInterface
  CreateBirthDay(w http.ResponseWriter, r *http.Request) Response.ApiResponseInterface
}

type BirthDayController struct {
  i_list_birth_day_interactor usecase_list_birth_day.ListBirthDayInteractorInterface
  i_create_birth_day_interactor usecase_create_birth_day.CreateBirthDayInteractorInterface
}

// construct
func NewBirthDayController(
  i_list_birth_day_interactor usecase_list_birth_day.ListBirthDayInteractorInterface,
  i_create_birth_day_interactor usecase_create_birth_day.CreateBirthDayInteractorInterface,
) BirthDayControllerInterface {
  return &BirthDayController {
    i_list_birth_day_interactor,
    i_create_birth_day_interactor,
  }
}

// list birth_day
func (controller_birthday BirthDayController) ListBirthDay(w http.ResponseWriter, r *http.Request) (Response.ApiResponseInterface) {
  user_id := domain_auth.GetUserIdByToken(w, r)
  return controller_birthday.i_list_birth_day_interactor.Handle(user_id)
}

// create birth_days
func (controller_birthday BirthDayController) CreateBirthDay(w http.ResponseWriter, r *http.Request) (Response.ApiResponseInterface) {
  // json decode
  input_data := []usecase_create_birth_day.CreateBirthDayRequest{}
  Request.JsonDecode(r, &input_data)
  user_id := domain_auth.GetUserIdByToken(w, r)
  input_data2 := []usecase_create_birth_day.CreateBirthDayRequest{}

  input_data2 = convertCreateInputData(input_data2, input_data, user_id)
  if errors := usecase_create_birth_day.IsValidRequestBody(input_data2); len(errors) > 0 {
    return Response.NewBadRequestResponse(errors)
  }
  return controller_birthday.i_create_birth_day_interactor.Handle(input_data2)
}

// Let the user_id bind to the object
func convertCreateInputData(
  result []usecase_create_birth_day.CreateBirthDayRequest,
  target []usecase_create_birth_day.CreateBirthDayRequest,
  user_id domain_user.UserId,
) ([]usecase_create_birth_day.CreateBirthDayRequest) {
  for _, v := range target {
    result = append(result, usecase_create_birth_day.CreateBirthDayRequest{
      UserId: user_id.GetValue(),
      Name: v.Name,
      Date: v.Date,
    })
  }
  return result
}