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
    birth_days := repository_birthday.ListBirthDay()
    return c.JSON(http.StatusOK, birth_days)
  }
}

func CreateBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    repository_birthday.CreateBirthDay(c)
    return c.JSON(http.StatusCreated, Empty.Empty{})
  }
}
