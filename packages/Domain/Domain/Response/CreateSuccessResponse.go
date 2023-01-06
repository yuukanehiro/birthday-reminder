package Response

import (
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type CreateSuccessResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

func NewCreateSuccessResponse() ApiResponseInterface {
  return &CreateSuccessResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: http.StatusCreated,
    Message: http.StatusText(http.StatusCreated),
    Errors: []Error{},
    Data: []Empty{},
  }
}

func (r CreateSuccessResponse) GetStatusCode() int {
  return r.StatusCode
}
