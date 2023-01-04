package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Empty"
  interactor_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  // "birthday-reminder/packages/Domain/Domain/Response"
)


func ListBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    interface_list_birth_day_interactor := interactor_birth_day.NewListBirthDayInteractor()
    birth_days := interface_list_birth_day_interactor.Handle()
    return c.JSON(http.StatusOK, birth_days)
  }
}

func CreateBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    interface_create_birth_day_interactor := interactor_birth_day.NewCreateBirthDayInteractor()
    interface_create_birth_day_interactor.Handle(c)
    return c.JSON(http.StatusCreated, Empty.Empty{})
  }
}
