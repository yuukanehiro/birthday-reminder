package main

import (
  "net/http"
  "os"
  "fmt"
)

func main() {
  err := http.ListenAndServe(
    ":8080",
    http.HandlerFunc(func(http_response_writer http.ResponseWriter, http_request *http.Request) {
      fmt.Fprintf(http_response_writer, "hello, %s", http_request.URL.Path[1:])
    }),
  )
  if err != nil {
    fmt.Printf("failed to terminate server: %v", err)
    os.Exit(1)
  }
}