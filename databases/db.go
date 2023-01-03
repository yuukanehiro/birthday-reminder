package databases

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
//   "os"
//   "time"
)

type Player struct {
  gorm.Model
  Name      string
  Age       int
  // Birthday  time.Time
  // Activated bool
  // Struct    Address
  // Pointer   *string
}


func InitDB() {
  dsn := "host=db user=docker password=docker dbname=birthday_reminder port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
      panic("failed to connect database")
  }
  db.AutoMigrate(&Player{})
  fmt.Println("migrated")
}

// func connectDB() {
// 	var err error
// 	db, err = gorm.Open(
// 		"postgres",
// 		fmt.Sprintf(
// 			"host=%s user=%s dbname=%s password=%s sslmode=disable",
// 			os.Getenv("POSTGRES_HOST"),
// 			os.Getenv("POSTGRES_USER"),
// 			os.Getenv("POSTGRES_DB"),
// 			os.Getenv("POSTGRES_PASSWORD"),
// 		),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }