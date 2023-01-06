package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
)

type ListBirthDayInteractor struct {
  i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface
}

// construct
func NewListBirthDayInteractor(
  i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface,
) usecase_list_birth_day.ListBirthDayInteractorInterface {
  return &ListBirthDayInteractor{i_birth_day_repo}
}

// execute
func (interactor_list_birth_day ListBirthDayInteractor) Handle() (birth_days []domain_birth_day.BirthDay) {
  birth_days = interactor_list_birth_day.i_birth_day_repo.ListBirthDay()
  return
}
