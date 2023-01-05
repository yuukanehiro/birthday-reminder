package BirthDay

import(
  "os"
  "log"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

type BirthDayRepository struct {}

func NewBirthDayRepository() domain_birth_day.BirthDayRepositoryInterface {
  return &BirthDayRepository{}
}


func (repository_birth_day BirthDayRepository) ListBirthDay() (birth_days []domain_birth_day.BirthDay) {
  // todo. 共通処理化
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  rdb := rdb_interface.ConnectDB()
  db, _ := rdb.DB()
  defer db.Close()

  birth_days = []domain_birth_day.BirthDay{}
  rdb.Select("id", "user_id", "date").Find(&birth_days)
  return
}

func (repository_birth_day BirthDayRepository) CreateBirthDay(dto usecase_create_birth_day.CreateBirthDayRequest) {
  // todo. 共通処理化
  rdb_interface := Rdb.NewRdbFactory(os.Getenv("DB_RDBMS"))
  rdb := rdb_interface.ConnectDB()
  db, _ := rdb.DB()
  defer db.Close()

  err := rdb.Exec("INSERT INTO birth_days (user_id, date) VALUES (?, ?)", dto.UserId, dto.Date)
  if err != nil {
    log.Print(err)
    return
  }
}