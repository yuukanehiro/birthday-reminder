package Rdb

import (
  "gorm.io/gorm"
)

type RdbInterface interface {
  ConnectDB() *gorm.DB
}
