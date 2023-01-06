package Response

import (
  "birthday-reminder/packages/Domain/Domain/Timestamp"
)

type ApiResponse struct {
  Timestamp  string `json:"timestamp"`
  StatusCode int `json:"status_code"`
  Message    string `json:"message"`
  Errors     interface{} `json:"errors"`
  Data       interface{} `json:"data"`
}

func NewApiResponse(
  status_code int,
  message string,
  errors interface{},
  data interface{},
) *ApiResponse {
  return &ApiResponse {
    Timestamp: Timestamp.GetNowTimeByISO8601Format(),
    StatusCode: 200,
    Message: message,
    Errors: errors,
    Data: data,
  }
}
