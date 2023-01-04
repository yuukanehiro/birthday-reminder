package BirthDay

type BirthDay struct {
  Id int `json:"id" param:"id"`
  UserId int `json:"user_id"`
  Date string `json:"date"`
}