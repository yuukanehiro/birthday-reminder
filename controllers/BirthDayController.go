package controllers

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  "log"
  "net/http"
  "encoding/json"
)

type BirthDayControllerInterface interface {
  ListBirthDay(w http.ResponseWriter, r *http.Request)
  CreateBirthDay(w http.ResponseWriter, r *http.Request)
}

type BirthDayController struct {
  i_list_birth_day_interactor usecase_list_birth_day.ListBirthDayInteractorInterface
  i_create_birth_day_interactor usecase_create_birth_day.CreateBirthDayInteractorInterface
}

func NewBirthDayController(
  i_list_birth_day_interactor usecase_list_birth_day.ListBirthDayInteractorInterface,
  i_create_birth_day_interactor usecase_create_birth_day.CreateBirthDayInteractorInterface,
) BirthDayControllerInterface {
  return &BirthDayController {
    i_list_birth_day_interactor,
    i_create_birth_day_interactor,
  }
}

func (controller_birthday BirthDayController) ListBirthDay(w http.ResponseWriter, r *http.Request) {
  birth_days := controller_birthday.i_list_birth_day_interactor.Handle()
  log.Print(birth_days)
  //return c.JSON(http.StatusOK, birth_days)

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
  output, _ := json.MarshalIndent(birth_days_response, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (controller_birthday BirthDayController) CreateBirthDay(w http.ResponseWriter, r *http.Request) {
  //return c.JSON(http.StatusCreated, controller_birthday.i_create_birth_day_interactor.Handle(c))
}
