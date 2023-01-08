package BirthDay

import (
  "github.com/golang-module/carbon"
)

type BirthDayAge struct {
  value int
}

type BirthDayAgeInterface interface {
  GetValue() int
}

// construct
func NewBirthDayAge(birth_day string) BirthDayAgeInterface {
  return BirthDayAge{
    value: calcAge(birth_day),
  }
}

func (b BirthDayAge) GetValue() int {
  return b.value
}

// calculate age
func calcAge(birth_day string) int {
  return carbon.Parse(birth_day).Age()
}
