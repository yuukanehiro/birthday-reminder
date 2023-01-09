package Date

import (
  "time"
  "gopkg.in/go-playground/validator.v9"
)

func IsDate(fl validator.FieldLevel) bool {
  _, err := time.Parse("2006-01-02", fl.Field().String())
  if err != nil {
    return false
  }
  return true
}
