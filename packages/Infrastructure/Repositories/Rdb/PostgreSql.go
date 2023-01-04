package Rdb

import (
  "fmt"
  "gorm.io/gorm"
  "os"
  "gorm.io/driver/postgres"
)

type PostgreSql struct {}

func (p PostgreSql) ConnectDB() *gorm.DB {
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
  )
  var err error
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  return db
}

