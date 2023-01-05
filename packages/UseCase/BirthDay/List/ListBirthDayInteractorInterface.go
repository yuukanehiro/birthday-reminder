package List

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
)

type ListBirthDayInteractorInterface interface {
  Handle() (birth_days []domain_birth_day.BirthDay)
}
