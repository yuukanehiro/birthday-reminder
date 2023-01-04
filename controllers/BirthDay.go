package controllers

import(
  "net/http"
  "os"
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  "birthday-reminder/packages/Domain/Domain/BirthDay"
  "birthday-reminder/packages/Domain/Domain/Empty"
)


func GetBirthDays() echo.HandlerFunc {
  return func(c echo.Context) error {
    rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
    db := rdb_interface.ConnectDB()
    birth_days := []BirthDay.BirthDay{}
    db.Find(&birth_days)
    return c.JSON(http.StatusOK, birth_days)
  }
}

func CreateBirthDay() echo.HandlerFunc {
  return func(c echo.Context) error {
    rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
    db := rdb_interface.ConnectDB()
    birth_day := BirthDay.BirthDay{
      Id: 102, UserId: 1, Date: "2022-01-01",
    }
    if err := c.Bind(&birth_day); err != nil {
      return err
    }
    db.Create(&birth_day)
    return c.JSON(http.StatusCreated, Empty.Empty{})
  }
}
