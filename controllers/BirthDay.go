package controllers

import(
  "net/http"
  "os"
  "log"
  "github.com/labstack/echo"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  "birthday-reminder/packages/Domain/Domain/BirthDay"
  "birthday-reminder/packages/Domain/Domain/Empty"
  // "birthday-reminder/packages/Domain/Domain/Response"
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
    birth_day := new(BirthDay.BirthDay)
    if err := c.Bind(birth_day)
    err != nil {
      log.Printf("err %v", err.Error())
      return c.String(http.StatusInternalServerError, "Error!")
    }
    db.Create(&birth_day)
    // todo. Responseに分離させる
    //return c.JSON(Response.NewResponse(http.StatusCreated, Empty.Empty{}))
    return c.JSON(http.StatusCreated, Empty.Empty{})
  }
}
