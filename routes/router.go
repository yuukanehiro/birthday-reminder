package routes

import (
  "net/http"
  "birthday-reminder/controllers"
)

type RouterInterface interface {
  HandleBirthDayRequest(w http.ResponseWriter, r *http.Request)
  HandleAuthRegisterRequest(w http.ResponseWriter, r *http.Request)
}

type Router struct {
  i_birth_day_controller controllers.BirthDayControllerInterface
  i_auth_register_controller controllers.AuthRegisterControllerInterface
}

// construct
func NewRouter(
  i_birth_day_controller controllers.BirthDayControllerInterface,
  i_auth_register_controller controllers.AuthRegisterControllerInterface,
) RouterInterface {
  return &Router{
    i_birth_day_controller,
    i_auth_register_controller,
  }
}
