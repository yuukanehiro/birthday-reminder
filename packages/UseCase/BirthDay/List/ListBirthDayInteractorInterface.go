package List

import (
  "birthday-reminder/packages/Domain/Domain/Response"
)

type ListBirthDayInteractorInterface interface {
  Handle() (Response.ApiResponseInterface)
}
