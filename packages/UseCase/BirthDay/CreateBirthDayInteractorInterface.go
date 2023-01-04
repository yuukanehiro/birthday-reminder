package BirthDay

import (
  "github.com/labstack/echo"
)

type CreateBirthDayInteractorInterface interface {
  Handle(c echo.Context)
}
