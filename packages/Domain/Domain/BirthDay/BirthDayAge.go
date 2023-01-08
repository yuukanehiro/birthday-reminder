package BirthDay

import (
  "github.com/golang-module/carbon"
)

type BirthDayAge struct {
  Value int
}

type BirthDayAgeInterface interface {
  GetValue() int
}

func NewBirthDayAge(birth_day string) BirthDayAgeInterface {
  return BirthDayAge{
    Value: carbon.Parse(birth_day).Age(),
  }
}

func (b BirthDayAge) GetValue() int {
  return b.Value
}
