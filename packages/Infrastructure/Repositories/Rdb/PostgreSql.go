package Rdb

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "birthday-reminder/config"
)

type PostgreSql struct {}

// construct
func NewPostgreSql() *PostgreSql {
  return &PostgreSql{}
}

// create gorm instance for PostgreSQL
func (p PostgreSql) ConnectDB() *gorm.DB {
  cfg := config.NewConfig()
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
    cfg.DB_HOST,
    cfg.DB_USER,
    cfg.DB_PASSWORD,
    cfg.DB_NAME,
    cfg.DB_PORT,
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database.")
  }
  return db
}
