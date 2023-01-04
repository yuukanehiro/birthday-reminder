package Response

type Response struct {
  Status  int         `json:"status"`
  Message string      `json:"message"`
  Result  interface{} `json:"result"`
}

func NewResponse(status int, message string, result interface{}) *Response {
  return &Response{status, message, result}
}
