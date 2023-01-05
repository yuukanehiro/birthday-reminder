package Create

type CreateBirthDayRequest struct {
  UserId int  `json:"user_id"`
  Date string `json:"date"`
}
