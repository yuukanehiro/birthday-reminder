package Response

import (
  "net/http"
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type GetSuccessResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

func NewGetSuccessResponse(data interface{}) ApiResponseInterface {
  return &GetSuccessResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: http.StatusOK,
    Message: "",
    Errors: []Error{},
    Data: data,
  }
}

func (r GetSuccessResponse) GetStatusCode() int {
	return r.StatusCode
  }
