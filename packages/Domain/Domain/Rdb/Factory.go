package Rdb

import (
  repository_rdb "birthday-reminder/packages/Infrastructure/Repositories/Rdb"
  // mock_repository_rdb "birthday-reminder/packages/MockInfrastructure/Repositories/Rdb"
)


func NewRdbFactory(rdbms_name string) RdbInterface {
  // todo. mapで動的に切り替えられるようにする。ifやswitchは絶対やめたい
  // if rdbms_name == "postgresql" {
  //   return &PostgreSql{}
  // }
  // return &MockDb{}
  return &repository_rdb.PostgreSql{}
}
