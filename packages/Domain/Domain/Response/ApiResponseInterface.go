package Response

type ApiResponseInterface interface {
  // GetStatusCode get HTTP status code
  GetStatusCode() int
}
