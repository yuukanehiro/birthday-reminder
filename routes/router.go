package routes

import (
  "net/http"
  "encoding/json"
  "birthday-reminder/controllers"
  "birthday-reminder/packages/Domain/Domain/Response"
)

type RouterInterface interface {
  HandleBirthDayRequest(w http.ResponseWriter, r *http.Request)
}

type Router struct {
  i_birth_day_controller controllers.BirthDayControllerInterface
}

// construct
func NewRouter(
  i_birth_day_controller controllers.BirthDayControllerInterface) RouterInterface {
  return &Router{i_birth_day_controller}
}

// Assign controller method by HTTP Request Method
func proxy(router Router, w http.ResponseWriter, r *http.Request) Response.ApiResponseInterface {
  switch r.Method {
    case "GET":
      return router.i_birth_day_controller.ListBirthDay(w, r)
    case "POST":
	  return router.i_birth_day_controller.CreateBirthDay(w, r)
    default:
      return router.i_birth_day_controller.ListBirthDay(w, r)
  }
}

func (router Router) HandleBirthDayRequest(w http.ResponseWriter, r *http.Request) {
  response := proxy(router, w, r)
  output, _ := json.MarshalIndent(response, "", "\t\t")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.GetStatusCode())
  w.Write(output)
}