package List

import (
  domain_birth_day "birthday-reminder/packages/Domain/Domain/BirthDay"
)


type BirthDayResponse struct {
  Id int      `json:"id" param:"id" gorm:"primarykey;unique;type:int"`
  UserId int  `json:"user_id"       gorm:"not null;type:int"`
  Name string `json:"name"          gorm:"not null;type:varchar(255)"`
  Age int     `json:"age"           gorm:"not null;type:int"`
  Date string `json:"date"          gorm:"not null;type:varchar(10)"`
}

// construct
func NewBirthDayResponse(
  id int,
  user_id int,
  name string,
  age domain_birth_day.BirthDayAge,
  date string,
) BirthDayResponse {
  return BirthDayResponse{
    Id: id,
    UserId: user_id,
    Name: name,
    Age: age.GetValue(),
    Date: date,
  }
}

type BirthDaysResponse struct {
  BirthDays []BirthDayResponse
}
