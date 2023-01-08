package List

type BirthDayResponse struct {
  Id int      `json:"id" param:"id" gorm:"primarykey;unique;type:int"`
  UserId int  `json:"user_id"       gorm:"not null;type:int"`
  Name string `json:"name"          gorm:"not null;type:varchar(255)"`
  Age int     `json:"age"           gorm:"not null;type:int"`
  Date string `json:"date"          gorm:"not null;type:varchar(10)"`
}

type BirthDaysResponse struct {
  BirthDays []BirthDayResponse
}