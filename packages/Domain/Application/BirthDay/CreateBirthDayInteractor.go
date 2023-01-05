package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

type CreateBirthDayInteractor struct {
  i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface
}

func NewCreateBirthDayInteractor(i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface) usecase_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{i_birth_day_repo}
}

func (interactor_create_birth_day CreateBirthDayInteractor) Handle() {
  //interactor_create_birth_day.i_birth_day_repo.CreateBirthDay(birth_day)
  //interactor_create_birth_day.i_birth_day_repo.CreateBirthDay()
}
