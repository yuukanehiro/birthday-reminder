package controllers

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  "birthday-reminder/packages/Domain/Domain/Response"
  "net/http"
  "encoding/json"
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

// create birth_day
func (controller_birthday BirthDayController) CreateBirthDay(w http.ResponseWriter, r *http.Request) (Response.ApiResponseInterface) {
  body := make([]byte, r.ContentLength)
  r.Body.Read(body)
  var createBirthDayRequest usecase_create_birth_day.CreateBirthDayRequest
  json.Unmarshal(body, &createBirthDayRequest)
  result := usecase_create_birth_day.CreateBirthDayRequest{UserId: createBirthDayRequest.UserId, Date: createBirthDayRequest.Date}

  controller_birthday.i_create_birth_day_interactor.Handle(result)
  return Response.NewCreateSuccessResponse()
}
