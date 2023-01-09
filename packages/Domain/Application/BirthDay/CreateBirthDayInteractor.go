package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  "birthday-reminder/packages/Domain/Domain/Response"
)

type CreateBirthDayInteractor struct {
  i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface
}

// construct
func NewCreateBirthDayInteractor(i_birth_day_repo domain_birth_day.BirthDayRepositoryInterface) usecase_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{i_birth_day_repo}
}

// execute
func (interactor_create_birth_day CreateBirthDayInteractor) Handle(birth_days_request []usecase_create_birth_day.CreateBirthDayRequest) (Response.ApiResponseInterface) {
  interactor_create_birth_day.i_birth_day_repo.CreateBirthDay(birth_days_request)
  return Response.NewCreateSuccessResponse()
}
