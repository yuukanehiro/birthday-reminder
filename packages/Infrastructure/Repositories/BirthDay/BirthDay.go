package BirthDay

import(
  "os"
  "github.com/labstack/echo"
  "log"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  "birthday-reminder/packages/Domain/Domain/BirthDay"
)

type BirthDayRepository struct {
}

func NewBirthDayRepository() BirthDay.BirthDayRepositoryInterface {
  return &BirthDayRepository{}
}


func (repository_birth_day BirthDayRepository) ListBirthDay() (birth_days []BirthDay.BirthDay) {
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  db := rdb_interface.ConnectDB()
  birth_days = []BirthDay.BirthDay{}
  db.Select("id", "user_id", "date").Find(&birth_days)
  return
}

func (repository_birth_day BirthDayRepository) CreateBirthDay(c echo.Context) {
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  db := rdb_interface.ConnectDB()
  birth_day := new(BirthDay.BirthDay)
  if err := c.Bind(birth_day)
  err != nil {
    log.Printf("err %v", err.Error())
  }
  db.Create(&birth_day)
}