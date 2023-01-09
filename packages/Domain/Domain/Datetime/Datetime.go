package Datetime

import(
  "github.com/golang-module/carbon"
)

// GetDatetime get datetime ex. 2023-01-09 17:29:22
func GetDatetime() string {
  return carbon.Now().ToDateTimeString()
}
