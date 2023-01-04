package BirthDay

import(
  "os"
  "github.com/labstack/echo"
  "log"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  "birthday-reminder/packages/Domain/Domain/BirthDay"
  "birthday-reminder/packages/Domain/Domain/Empty"
)

type BirthDayRepository struct {
}

func NewBirthDayRepository() BirthDay.BirthDayRepositoryInterface {
  return &BirthDayRepository{}
}


func (repository_birth_day BirthDayRepository) ListBirthDay() (birth_days []BirthDay.BirthDay) {
  // todo. 共通処理化
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  rdb := rdb_interface.ConnectDB()
  db, _ := rdb.DB()
  defer db.Close()

  birth_days = []BirthDay.BirthDay{}
  rdb.Select("id", "user_id", "date").Find(&birth_days)
  return
}

func (repository_birth_day BirthDayRepository) CreateBirthDay(c echo.Context) *Empty.Empty {
  // todo. 共通処理化
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  rdb := rdb_interface.ConnectDB()
  db, _ := rdb.DB()
  defer db.Close()

  birth_day := new(BirthDay.BirthDay)
  if err := c.Bind(birth_day)
  err != nil {
    log.Printf("err %v", err.Error())
  }
  rdb.Create(&birth_day)
  return &Empty.Empty{}
}