package Rdb

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "birthday-reminder/config"
  domain_rdb "birthday-reminder/packages/Domain/Domain/Rdb"
)

type Rdb struct {
  cfg *config.Config
}

// construct
func NewRdb(cfg *config.Config) domain_rdb.RdbInterface {
  return &Rdb{cfg: cfg}
}

// create gorm instance for PostgreSQL
func (r Rdb) ConnectDB() *gorm.DB {
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
    r.cfg.DB_HOST,
    r.cfg.DB_USER,
    r.cfg.DB_PASSWORD,
    r.cfg.DB_NAME,
    r.cfg.DB_PORT,
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database.")
  }
  return db
}
