package Response

import (
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type MethodNotAllowedResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

// construct
func NewMethodNotAllowedResponse() ApiResponseInterface {
  return &MethodNotAllowedResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: http.StatusMethodNotAllowed,
    Message: http.StatusText(http.StatusMethodNotAllowed),
    Errors: []Error{},
    Data: []Empty{},
  }
}

func (r MethodNotAllowedResponse) GetStatusCode() int {
  return r.StatusCode
}
