package Auth

import (
  "birthday-reminder/packages/Domain/Domain/Response"
  domain_auth "birthday-reminder/packages/Domain/Domain/Auth"
  usecase_auth_register "birthday-reminder/packages/UseCase/Auth/Register"
)

type AuthRegisterInteractor struct {
  i_auth_register_repo domain_auth.AuthRegisterRepositoryInterface
}

// construct
func NewAuthRegisterInteractor(
  i_auth_register_repo domain_auth.AuthRegisterRepositoryInterface,
) usecase_auth_register.AuthRegisterInteractorInterface {
  return &AuthRegisterInteractor{i_auth_register_repo}
}

// execute
func (interactor AuthRegisterInteractor) Handle() Response.ApiResponseInterface {
  user_id := interactor.i_auth_register_repo.AuthRegister()
  jwt_token_string := domain_auth.GetTokenByUserId(user_id)
  response_auth_register := usecase_auth_register.NewAuthRegisterResponse(jwt_token_string)
  return Response.NewCreateSuccessResponseWithData(response_auth_register)
}
