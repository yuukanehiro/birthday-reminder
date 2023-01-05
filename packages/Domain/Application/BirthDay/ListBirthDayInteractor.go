package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  repository_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)

type ListBirthDayInteractor struct {}

func NewListBirthDayInteractor() usecase_list_birth_day.ListBirthDayInteractorInterface {
  return &ListBirthDayInteractor{}
}

func (interactor_list_birth_day ListBirthDayInteractor) Handle() (birth_days []domain_birth_day.BirthDay) {
  interface_birth_day := repository_birth_day.NewBirthDayRepository()
  birth_days = interface_birth_day.ListBirthDay()
  return
}
