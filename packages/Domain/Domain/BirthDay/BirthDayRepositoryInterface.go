package BirthDay

import (
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

type BirthDayRepositoryInterface interface {
  ListBirthDay() (birth_days []BirthDay)
  CreateBirthDay(birth_days_request []usecase_create_birth_day.CreateBirthDayRequest)
}
