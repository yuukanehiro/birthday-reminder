package controllers

import(
  "net/http"
  "os"
  "github.com/labstack/echo"
  "birthday-reminder/packages/Infrastructure/Repositories/Rdb"
)

type BirthDay struct {
  Id int `json:"id" param:"id"`
  UserId int `json:"user_id"`
  Date string `json:"date"`
}

func GetBirthDays() echo.HandlerFunc {
  return func(c echo.Context) error {
    rdb_interface := Rdb.NewRdb(os.Getenv("DB_RDBMS"))
    db := rdb_interface.ConnectDB()
    birth_days := []BirthDay{}
    db.Find(&birth_days)
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