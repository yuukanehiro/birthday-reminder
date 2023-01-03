package main

import (
  "birthday-reminder/routes"
  "birthday-reminder/packages/Infrastructure/Repositories/Rdb"
  "birthday-reminder/packages/Domain/Domain"
)

var (
	rdb_interface Domain.RdbInterface
)

func main() {
  e := routes.Init()
  rdb_interface = Rdb.NewRdb()
  rdb_interface.ConnectDB()

  e.Logger.Fatal(e.Start(":8080"))
}
