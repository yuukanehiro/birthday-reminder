package routes

import (
  "birthday-reminder/controllers"
  "github.com/labstack/echo"
)

func Init() *echo.Echo {
  e := echo.New()
  api := e.Group("/api/v1")
  {
    api.GET("/birth-days", controllers.GetBirthDays())
    api.POST("/birth-days", controllers.PostBirthDays())
  }
  return e
}