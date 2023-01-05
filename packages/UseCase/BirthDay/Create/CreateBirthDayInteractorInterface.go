package Create

import (
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Empty"
)

type CreateBirthDayInteractorInterface interface {
  Handle(c echo.Context) *Empty.Empty
}
