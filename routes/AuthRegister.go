package routes

import (
  "net/http"
  "encoding/json"
  "birthday-reminder/packages/Domain/Domain/Response"
)

// Assign controller method by HTTP Request Method
func proxyAuthRegister(router Router, w http.ResponseWriter, r *http.Request) Response.ApiResponseInterface {
  switch r.Method {
    case "POST":
      return router.i_auth_register_controller.AuthRegister()
    default:
      return Response.NewMethodNotAllowedResponse()
  }
}

func (router Router) HandleAuthRegisterRequest(w http.ResponseWriter, r *http.Request) {
  response := proxyAuthRegister(router, w, r)
  output, _ := json.MarshalIndent(response, "", "\t\t")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.GetStatusCode())
  w.Write(output)
}
