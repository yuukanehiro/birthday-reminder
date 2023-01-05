package controllers

import (
	usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
	usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
	"log"
	"net/http"
)

type BirthDayControllerInterface interface {
  ListBirthDay()
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

func (controller_birthday BirthDayController) ListBirthDay() {
  birth_days := controller_birthday.i_list_birth_day_interactor.Handle()
  log.Print(birth_days)
  //return c.JSON(http.StatusOK, birth_days)
}

func (controller_birthday BirthDayController) CreateBirthDay(w http.ResponseWriter, r *http.Request) {
  //return c.JSON(http.StatusCreated, controller_birthday.i_create_birth_day_interactor.Handle(c))
}
