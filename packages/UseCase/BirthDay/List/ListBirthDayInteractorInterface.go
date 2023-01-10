package List

import (
  "birthday-reminder/packages/Domain/Domain/Response"
  domain_user "birthday-reminder/packages/Domain/Domain/User"
)

type ListBirthDayInteractorInterface interface {
  Handle(user_id domain_user.UserId) (Response.ApiResponseInterface)
}
