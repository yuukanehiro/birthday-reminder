package BirthDay

import(
  "log"
  "gorm.io/gorm"
  "birthday-reminder/config"
  "birthday-reminder/packages/Domain/Domain/Rdb"
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

var (
  rdb *gorm.DB
)

type BirthDayRepository struct {
}

func NewBirthDayRepository() domain_birth_day.BirthDayRepositoryInterface {
  connectDB()
  return &BirthDayRepository{}
}

func connectDB() {
  cfg, err := config.NewConfig()
  if err != nil {
    log.Fatalf("failed load config. %v", err)
  }
  rdb_interface := Rdb.NewRdbFactory(cfg.DB_RDBMS)
  rdb = rdb_interface.ConnectDB()
}

func (repository_birth_day BirthDayRepository) ListBirthDay() (birth_days []domain_birth_day.BirthDay) {
  db, _ := rdb.DB()
  defer db.Close()

  birth_days = []domain_birth_day.BirthDay{}
  rdb.Select("id", "user_id", "date").Find(&birth_days)
  return
}

func (repository_birth_day BirthDayRepository) CreateBirthDay(dto usecase_create_birth_day.CreateBirthDayRequest) {
  db, _ := rdb.DB()
  defer db.Close()

  rdb.Exec("INSERT INTO birth_days (user_id, date) VALUES (?, ?)", dto.UserId, dto.Date)
}