package Auth

import (
  "time"
  "github.com/dgrijalva/jwt-go"
  "birthday-reminder/config"
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
  cfg := config.NewConfig()
  claims := jwt.MapClaims{
    "user_id": user_id,
    "exp": time.Now().Add(time.Hour * time.Duration(cfg.JWT_TOKEN_EXPIRE_HOUR)).Unix(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  // Add Signature to Token
  tokenString, _ := token.SignedString([]byte(cfg.JWT_SECRET_KEY))
  response_auth_register := usecase_auth_register.NewAuthRegisterResponse(tokenString)
  return Response.NewCreateSuccessResponseWithData(response_auth_register)
}
