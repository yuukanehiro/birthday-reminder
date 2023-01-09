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
func (repository BirthDayRepository) ListBirthDay() (birth_days []domain_birth_day.BirthDay) {
  birth_days = []domain_birth_day.BirthDay{}
  if err := repository.rdb.Select("id, user_id, name, date").Find(&birth_days).Error; err != nil {
    panic("Error Select.")
  }
  return
}

// create birth_day
func (repository BirthDayRepository) CreateBirthDay(request []usecase_create_birth_day.CreateBirthDayRequest) {
  for _, v := range request {
    if err := repository.rdb.Exec(
      "INSERT INTO birth_days (user_id, name, date) VALUES (?, ?, ?)",
      v.UserId,
      v.Name,
      v.Date,
    ).Error; err != nil {
      panic("Error Insert.")
    }
  }
}
