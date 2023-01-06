package BirthDay

import(
  "gorm.io/gorm"
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
  usecase_create_birth_day "birthday-reminder/packages/UseCase/BirthDay/Create"
)

type BirthDayRepository struct {
  rdb *gorm.DB
}

// construct
func NewBirthDayRepository(rdb *gorm.DB) domain_birth_day.BirthDayRepositoryInterface {
  return &BirthDayRepository{
    rdb: rdb,
  }
}

// get list birth_days
func (repository_birth_day BirthDayRepository) ListBirthDay() (birth_days []domain_birth_day.BirthDay) {
  birth_days = []domain_birth_day.BirthDay{}
  repository_birth_day.rdb.Select("id", "user_id", "date").Find(&birth_days)
  return
}

// create birth_day
func (repository_birth_day BirthDayRepository) CreateBirthDay(dto usecase_create_birth_day.CreateBirthDayRequest) {
  repository_birth_day.rdb.Exec("INSERT INTO birth_days (user_id, date) VALUES (?, ?)", dto.UserId, dto.Date)
}