package Response

import (
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type BadRequestResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

func NewBadRequestResponse(errors []Error) ApiResponseInterface {
  return &BadRequestResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: http.StatusBadRequest,
    Message: http.StatusText(http.StatusBadRequest),
    Errors: errors,
    Data: []Empty{},
  }
}

func (r BadRequestResponse) GetStatusCode() int {
  return r.StatusCode
}
