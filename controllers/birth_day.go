package controllers

import(
  "net/http"
  "github.com/labstack/echo"
)

type BirthDay struct {
  Id int `json:"id"`
  Date string `json:"date"`
}

func GetBirthDays() echo.HandlerFunc {
  return func(c echo.Context) error {
    birth_days := []BirthDay {
      {Id: 1, Date: "1984 11/21"},
      {Id: 2, Date: "1984 11/22"},
	}
    return c.JSON(http.StatusOK, birth_days)
  }
}

func PostBirthDays() echo.HandlerFunc {
  return func(c echo.Context) error {
    birth_days := []BirthDay {
      {Id: 1, Date: "1984 11/21"},
      {Id: 2, Date: "1984 11/22"},
    }
    return c.JSON(http.StatusOK, birth_days)
  }
}