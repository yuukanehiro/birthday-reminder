package Rdb

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "birthday-reminder/config"
  domain_rdb "birthday-reminder/packages/Domain/Domain/Rdb"
)

type Rdb struct {
  i_rdb domain_rdb.RdbInterface
}

// construct
func NewRdb() domain_rdb.RdbInterface {
  return &Rdb{}
}

// create gorm instance for PostgreSQL
func (r Rdb) ConnectDB(cfg *config.Config) *gorm.DB {
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
