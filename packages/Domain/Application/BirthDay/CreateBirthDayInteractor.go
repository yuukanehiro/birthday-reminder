package BirthDay

import (
  "github.com/labstack/echo"
  interface_create_birth_day "birthday-reminder/packages/UseCase/BirthDay"
  repository_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)

type CreateBirthDayInteractor struct {}

func NewCreateBirthDayInteractor() interface_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{}
}

func (interactor_create_birth_day CreateBirthDayInteractor) Handle(c echo.Context) {
  interface_birth_day := repository_birth_day.NewBirthDayRepository()
  interface_birth_day.CreateBirthDay(c)
}
