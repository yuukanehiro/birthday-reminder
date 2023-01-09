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
func NewCreateBirthDayInteractor(i_repository domain_birth_day.BirthDayRepositoryInterface) usecase_create_birth_day.CreateBirthDayInteractorInterface {
  return &CreateBirthDayInteractor{i_repository}
}

// execute
func (interactor CreateBirthDayInteractor) Handle(request []usecase_create_birth_day.CreateBirthDayRequest) (Response.ApiResponseInterface) {
  interactor.i_birth_day_repo.CreateBirthDay(request)
  return Response.NewCreateSuccessResponse()
}
