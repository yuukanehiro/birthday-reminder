package List

import (
  "birthday-reminder/packages/Domain/Domain/Response"
)

type ListBirthDayInteractorInterface interface {
  Handle(user_id int64) (Response.ApiResponseInterface)
}
