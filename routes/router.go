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
func proxy(router Router, w http.ResponseWriter, r *http.Request) interface{} {
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
  result := proxy(router, w, r)
  errors := []Response.Error{}
  response := Response.NewApiResponse(200, "", errors, result)
  output, _ := json.MarshalIndent(response, "", "\t\t")

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(output)
}