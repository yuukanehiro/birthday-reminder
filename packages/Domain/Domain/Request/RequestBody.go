package Request

import (
  "net/http"
  "encoding/json"
  "io"
  "log"
)

// Map request body to specified object
func JsonDecode[T any](r *http.Request, inputData *T) {
  body, err := io.ReadAll(r.Body)
  if err != nil {
    log.Fatalf("failed JsonDecode when io.ReadAll. %v", err)
  }
  // finally must close
  defer r.Body.Close()
  err = json.Unmarshal(body, &inputData)
  if err != nil {
    log.Fatalf("failed JsonDecode when Unmarshal %v", err)
  }
}
