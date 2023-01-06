package Create

type CreateBirthDayInteractorInterface interface {
  Handle(birth_days_request []CreateBirthDayRequest)
}
