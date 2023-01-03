package main

import (
  "birthday-reminder/routes"
  "birthday-reminder/databases/rdb"
)


func main() {
  e := routes.Init()
  rdb.ConnectDB()

  e.Logger.Fatal(e.Start(":8080"))
}
