package controllers

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
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
  birth_days := controller_birthday.i_list_birth_day_interactor.Handle()

  birth_days_response := []usecase_list_birth_day.BirthDayResponse{}
  for _, v := range birth_days {
    birth_days_response = append(
      birth_days_response,
      usecase_list_birth_day.BirthDayResponse{
        Id: v.Id,
        UserId: v.UserId,
        Date: v.Date,
      },
    )
  }
  return Response.NewGetSuccessResponse(birth_days_response)
}

// create birth_days
func (controller_birthday BirthDayController) CreateBirthDay(w http.ResponseWriter, r *http.Request) (Response.ApiResponseInterface) {
  input_data := []usecase_create_birth_day.CreateBirthDayRequest{}
  Request.JsonDecode(r, &input_data)
  var data []usecase_create_birth_day.CreateBirthDayRequestValidate
  for _, v := range input_data {
    data = append(data, usecase_create_birth_day.CreateBirthDayRequestValidate{UserId: v.UserId, Date: v.Date,})
  }
  errors := usecase_create_birth_day.IsValid(data)
  if len(errors) > 0 {
    return Response.NewBadRequestResponse(errors)
  }
  controller_birthday.i_create_birth_day_interactor.Handle(input_data)
  return Response.NewCreateSuccessResponse()
}
