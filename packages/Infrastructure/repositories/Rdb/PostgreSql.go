package Rdb

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "os"
  "birthday-reminder/packages/Domain/Domain"
)

var DB *gorm.DB

type PostgreSql struct {}
type MockDb struct {}

func NewRdb(rdbms_name string) Domain.RdbInterface {
  return &PostgreSql{}
  // todo. mapで動的に切り替えられるようにする
  // if rdbms_name == "postgresql" {
  //   return &PostgreSql{}
  // }
  // return &MockDb{}
}

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

func (m MockDb) ConnectDB() {
  //
}