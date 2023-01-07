package Date

import (
  "time"
  "gopkg.in/go-playground/validator.v9"
  "log"
)

func IsDate(fl validator.FieldLevel) bool {
  _, err := time.Parse("2006-01-02", fl.Field().String())
  log.Println("IsDate run start")
  if err != nil {
	log.Println(fl.Field().String())
	log.Println("IsDate run ... false")
	log.Println("IsDate run end")
    return false
  }
  log.Println("IsDate run ... true")
  log.Println("IsDate run end")
  return true
}
