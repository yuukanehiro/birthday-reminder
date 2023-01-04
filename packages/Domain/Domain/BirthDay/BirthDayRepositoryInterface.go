package BirthDay

import(
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Empty"
)

type BirthDayRepositoryInterface interface {
  ListBirthDay() (birth_days []BirthDay)
  CreateBirthDay(c echo.Context) *Empty.Empty
}
