package Response

type Error struct {
  message string
  property string
}

// func NewResponse(status int, message string, result interface{}) *Response {
//   return &Response{status, message, result}
// }
