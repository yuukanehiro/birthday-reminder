package BirthDay

import(
  "github.com/labstack/echo"
)

type BirthDayRepositoryInterface interface {
  ListBirthDay() (birth_days []BirthDay)
  CreateBirthDay(c echo.Context)
}
