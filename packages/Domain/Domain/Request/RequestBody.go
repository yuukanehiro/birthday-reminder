package Request

import (
  "net/http"
  "encoding/json"
  "io"
  "fmt"
)

// Map request body to specified object
func JsonDecode[T any](r *http.Request, inputData *T) {
  body, err := io.ReadAll(r.Body)
  if err != nil {
  	fmt.Println(err)
  }
  // 最後は閉じる
  defer r.Body.Close()
  err = json.Unmarshal(body, &inputData)
  if err != nil {
  	fmt.Println(err)
  }
}
