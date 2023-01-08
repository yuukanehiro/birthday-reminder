package BirthDay

import (
  "github.com/golang-module/carbon"
)

type BirthDayAge struct {
  Value int
}

func CalcAge(birth_day string) int {
  return carbon.Parse(birth_day).Age()
}
