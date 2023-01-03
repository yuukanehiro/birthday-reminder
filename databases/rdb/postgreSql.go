package rdb

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "time"
  "os"
)

type Player struct {
  gorm.Model
  Name      string
  Age       int
  Birthday  time.Time
}


func ConnectDB() {
  dsn := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB"),
	os.Getenv("POSTGRES_PORT"),
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(&Player{})
  fmt.Println("migrated")
}