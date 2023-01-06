package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

type CreateBirthDayInteractor struct {
  i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface
}

// construct
func NewCreateBirthDayInteractor(i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface) usecase_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{i_birth_day_repo}
}

// execute
func (interactor_create_birth_day CreateBirthDayInteractor) Handle(dto usecase_create_birth_day.CreateBirthDayRequest) {
  interactor_create_birth_day.i_birth_day_repo.CreateBirthDay(dto)
}
