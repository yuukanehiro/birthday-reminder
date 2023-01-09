package Auth

import(
  "gorm.io/gorm"
  domain_auth "birthday-reminder/packages/Domain/Domain/Auth"
  "birthday-reminder/packages/Domain/Domain/Datetime"
)

type AuthRegisterRepository struct {
  rdb *gorm.DB
}

// construct
func NewAuthRegisterRepository(rdb *gorm.DB) domain_auth.AuthRegisterRepositoryInterface {
  return &AuthRegisterRepository{
    rdb: rdb,
  }
}

type User struct {
  Id int `gorm:"primaryKey"`
  InsertDate string
}

// AuthRegister register user
func (repo AuthRegisterRepository) AuthRegister() int {
  now_datetime := Datetime.GetDatetime()
  user1 := User{InsertDate: now_datetime}

  if err := repo.rdb.Create(&user1).Error; err != nil {
    panic("Error Create. AuthRegister()")
  }
  return user1.Id
}
