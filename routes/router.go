package routes

import (
  "net/http"
  "birthday-reminder/controllers"
)

type RouterInterface interface {
  HandleBirthDayRequest(w http.ResponseWriter, r *http.Request)
}

type Router struct {
  i_birth_day_controller controllers.BirthDayControllerInterface
}

func NewRouter(
  i_birth_day_controller controllers.BirthDayControllerInterface) RouterInterface {
  return &Router{i_birth_day_controller}
}

func (router Router) HandleBirthDayRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
	router.i_birth_day_controller.ListBirthDay(w, r)
  case "POST":
	router.i_birth_day_controller.CreateBirthDay(w, r)
  default:
  	w.WriteHeader(405)
  }
}





