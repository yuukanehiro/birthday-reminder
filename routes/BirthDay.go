package routes

import (
  "net/http"
  "encoding/json"
  "birthday-reminder/packages/Domain/Domain/Response"
)

// Assign controller method by HTTP Request Method
func proxyBirthDay(router Router, w http.ResponseWriter, r *http.Request) Response.ApiResponseInterface {
  switch r.Method {
    case "GET":
      return router.i_birth_day_controller.ListBirthDay(w, r)
    case "POST":
      return router.i_birth_day_controller.CreateBirthDay(w, r)
    default:
      return Response.NewMethodNotAllowedResponse()
  }
}

func (router Router) HandleBirthDayRequest(w http.ResponseWriter, r *http.Request) {
  response := proxyBirthDay(router, w, r)
  output, _ := json.MarshalIndent(response, "", "\t\t")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.GetStatusCode())
  w.Write(output)
}