package Rdb

import (
  "gorm.io/gorm"
  "birthday-reminder/config"
)

type RdbInterface interface {
  ConnectDB(cfg *config.Config) *gorm.DB
}
