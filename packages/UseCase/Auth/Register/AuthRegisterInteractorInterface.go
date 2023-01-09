package Register

import (
  "birthday-reminder/packages/Domain/Domain/Response"
)

type AuthRegisterInteractorInterface interface {
  Handle() Response.ApiResponseInterface
}
