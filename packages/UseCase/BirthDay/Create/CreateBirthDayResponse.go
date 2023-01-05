package Create

type BirthDayResponse struct {
  Id *int     `json:"id" param:"id" gorm:"primarykey;unique;type:int"`
  UserId int  `json:"user_id"       gorm:"not null;type:int"`
  Date string `json:"date"          gorm:"not null;type:varchar(10)"`
}

type BirthDaysResponse struct {
  BirthDays []BirthDayResponse
}