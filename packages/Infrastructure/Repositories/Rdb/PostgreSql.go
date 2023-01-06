package Rdb

import (
  "fmt"
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "birthday-reminder/config"
)

type PostgreSql struct {}

func NewPostgreSql() *PostgreSql {
  return &PostgreSql{}
}

func (p PostgreSql) ConnectDB() *gorm.DB {
  cfg, err := config.NewConfig()
  if err != nil {
    log.Fatalf("failed load config.")
  }
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
