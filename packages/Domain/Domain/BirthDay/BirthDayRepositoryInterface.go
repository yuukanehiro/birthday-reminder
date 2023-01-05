package BirthDay

type BirthDayRepositoryInterface interface {
  ListBirthDay() (birth_days []BirthDay)
  //CreateBirthDay(BirthDay)
  CreateBirthDay()
}
