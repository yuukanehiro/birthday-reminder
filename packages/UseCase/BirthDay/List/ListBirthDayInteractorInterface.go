package List

type ListBirthDayInteractorInterface interface {
  Handle() (birth_days []BirthDayResponse)
}
