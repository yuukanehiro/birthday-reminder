package controllers

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  domain_auth "birthday-reminder/packages/Domain/Domain/Auth"
  "birthday-reminder/packages/Domain/Domain/Response"
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
  req, errors := usecase_create_birth_day.GetRequestBody(w, r)
  if errors != nil {
    _ = req
    return Response.NewBadRequestResponse(errors)
  }
  return controller_birthday.i_create_birth_day_interactor.Handle(req)
}
