package BirthDay

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
  domain_user "birthday-reminder/packages/Domain/Domain/User"
)

type BirthDayRepositoryInterface interface {
  ListBirthDay(user_id domain_user.UserId) (birth_days []BirthDay)
  CreateBirthDay(birth_days_request []usecase_create_birth_day.CreateBirthDayRequest)
}
