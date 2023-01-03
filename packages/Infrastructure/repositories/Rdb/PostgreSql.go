package Rdb

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "os"
  "birthday-reminder/packages/Domain/Domain"
)

type postgresql struct {}

func NewRdb() Domain.RdbInterface {
  return &postgresql{}
}

func (p postgresql) ConnectDB() {
  dsn := fmt.Sprintf(
	  "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	  os.Getenv("DB_HOST"),
	  os.Getenv("DB_USER"),
	  os.Getenv("DB_PASSWORD"),
	  os.Getenv("DB_NAME"),
	  os.Getenv("DB_PORT"),
  )
  _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
}