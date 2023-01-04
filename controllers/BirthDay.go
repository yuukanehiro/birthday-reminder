package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Empty"
  repository_birthday "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
  // "birthday-reminder/packages/Domain/Domain/Response"
)


func ListBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    interface_birth_day := repository_birthday.NewBirthDayRepository()
    birth_days := interface_birth_day.ListBirthDay()
    return c.JSON(http.StatusOK, birth_days)
  }
}

func CreateBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    interface_birth_day := repository_birthday.NewBirthDayRepository()
    interface_birth_day.CreateBirthDay(c)
    return c.JSON(http.StatusCreated, Empty.Empty{})
  }
}
