package main

import (
  "net/http"
  "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    api := e.Group("/api/v1")
    {
      api.GET("/birth-days", func(c echo.Context) error {
        return c.String(http.StatusOK, "birth days GET")
      })
      api.POST("/birth-days", func(c echo.Context) error {
        return c.String(http.StatusOK, "birth days POST")
      })
    }
    e.Logger.Fatal(e.Start(":8080"))
}
