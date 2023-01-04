package routes

import (
  "birthday-reminder/controllers"
  "github.com/labstack/echo"
)

func Init() *echo.Echo {
  e := echo.New()
  api := e.Group("/api/v1")
  {
    api.POST("/birth-day", controllers.CreateBirthDay())
    api.GET("/birth-days", controllers.ListBirthDay())
  }
  return e
}