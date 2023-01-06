package routes

import (
  "net/http"
  "encoding/json"
  "time"
  "birthday-reminder/controllers"
  "birthday-reminder/packages/Domain/Domain/Response"
)
const (
  TIMESTAMP_FORMAT_RFC3339 = "2006-01-02T15:04:05Z07:00"
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

// 共通APIレスポンス
type CommonResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

func handle(router Router, w http.ResponseWriter, r *http.Request) interface{} {
  switch r.Method {
    case "GET":
      return router.i_birth_day_controller.ListBirthDay(w, r)
    case "POST":
	  return router.i_birth_day_controller.CreateBirthDay(w, r)
    default:
      return router.i_birth_day_controller.ListBirthDay(w, r)
  }
}
func timeToString(t time.Time) string {
  str := t.Format(TIMESTAMP_FORMAT_RFC3339)
  return str
}

func (router Router) HandleBirthDayRequest(w http.ResponseWriter, r *http.Request) {
  result := handle(router, w, r)
  errors := []Response.Error{}
  response := &CommonResponse{
    Timestamp: timeToString(time.Now()),
	StatusCode: 200,
	Message: "",
	Errors: errors,
	Data: result,
  }
  output, _ := json.MarshalIndent(response, "", "\t\t")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(output)
}
