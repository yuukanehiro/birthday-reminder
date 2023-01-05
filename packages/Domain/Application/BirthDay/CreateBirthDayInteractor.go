package BirthDay

import (
  "github.com/labstack/echo"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  repository_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
  "birthday-reminder/packages/Domain/Domain/Empty"
)

type CreateBirthDayInteractor struct {}

func NewCreateBirthDayInteractor() usecase_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{}
}

func (interactor_create_birth_day CreateBirthDayInteractor) Handle(c echo.Context) *Empty.Empty {
  interface_birth_day := repository_birth_day.NewBirthDayRepository()
  return interface_birth_day.CreateBirthDay(c)
}
