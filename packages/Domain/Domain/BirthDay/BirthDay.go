package BirthDay

type BirthDay struct {
  Id int      `validate:"required,min=1"`
  UserId int  `validate:"required,min=1"`
  Date string `validate:"required,is_date"`
}
