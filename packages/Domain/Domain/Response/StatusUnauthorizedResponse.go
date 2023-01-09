package Response

import (
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type StatusUnauthorizedResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

// construct
func NewStatusUnauthorizedResponse() ApiResponseInterface {
  return &MethodNotAllowedResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: http.StatusUnauthorized,
    Message: http.StatusText(http.StatusUnauthorized),
    Errors: []Error{},
    Data: []Empty{},
  }
}

func (r StatusUnauthorizedResponse) GetStatusCode() int {
  return r.StatusCode
}
