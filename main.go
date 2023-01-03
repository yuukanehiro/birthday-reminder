package main

import (
  "birthday-reminder/routes"
  "birthday-reminder/databases"
)


func main() {
  e := routes.Init()

  databases.InitDB()

  e.Logger.Fatal(e.Start(":8080"))
}
