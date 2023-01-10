package BirthDay

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_list_birth_day "birthday-reminder/packages/UseCase/BirthDay/List"
  domain_user "birthday-reminder/packages/Domain/Domain/User"
  "birthday-reminder/packages/Domain/Domain/Response"
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
func (interactor ListBirthDayInteractor) Handle(user_id domain_user.UserId) (Response.ApiResponseInterface) {
  birth_days := interactor.i_birth_day_repo.ListBirthDay(user_id)
  birth_days_response := []usecase_list_birth_day.BirthDayResponse{}
  for _, v := range birth_days {
    vo_age := domain_birth_day.NewBirthDayAge(v.Date)

    birth_days_response = append(
      birth_days_response,
      usecase_list_birth_day.NewBirthDayResponse(
        v.Id,
        v.UserId,
        v.Name,
        vo_age,
        v.Date,
      ),
    )
  }
  return Response.NewGetSuccessResponse(birth_days_response)
}
