package Create

import (
  "birthday-reminder/packages/Domain/Domain/Response"
)

type CreateBirthDayInteractorInterface interface {
  Handle(birth_days_request []CreateBirthDayRequest) (Response.ApiResponseInterface)
}
