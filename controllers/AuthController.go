package controllers

type AuthController struct {}

type AuthControllerInterface interface {
}

func NewAuthController() AuthControllerInterface {
  return AuthController{}
}

func (c AuthController) Register() {
  
}