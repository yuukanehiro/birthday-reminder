package BirthDay

type BirthDay struct {
  Id int      `json:"id" param:"id" gorm:"unique;not null;type:int"`
  UserId int  `json:"user_id"       gorm:"not null;type:int"`
  Date string `json:"date"          gorm:"not null;type:varchar(10)"`
}