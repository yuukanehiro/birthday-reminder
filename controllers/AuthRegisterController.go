package controllers

import (
  "birthday-reminder/packages/Domain/Domain/Response"
  usecase_auth_register "birthday-reminder/packages/UseCase/Auth/Register"
)

type AuthRegisterController struct {
  interactor_auth_register usecase_auth_register.AuthRegisterInteractorInterface
}

type AuthRegisterControllerInterface interface {
  AuthRegister() Response.ApiResponseInterface
}

// construct
func NewAuthRegisterController(interactor_auth_register usecase_auth_register.AuthRegisterInteractorInterface) AuthRegisterControllerInterface {
  return &AuthRegisterController{
    interactor_auth_register: interactor_auth_register,
  }
}

// AuthRegister register users table and return access_token
func (c AuthRegisterController) AuthRegister() Response.ApiResponseInterface {
  return c.interactor_auth_register.Handle()
}
