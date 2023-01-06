package main

import (
  "fmt"
  "net/http"
  "gorm.io/gorm"
  "birthday-reminder/config"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  domain_rdb "birthday-reminder/packages/Domain/Domain/Rdb"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  infra_repo_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)

var cfg = config.NewConfig()
var rdb = newRDB(cfg)
var repo_birth_day = infra_repo_birth_day.NewBirthDayRepository(rdb)
var controller_birthday = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(repo_birth_day),
  app_birth_day.NewCreateBirthDayInteractor(repo_birth_day),
)
var router = routes.NewRouter(controller_birthday)


func main() {
  server := http.Server{
    Addr: fmt.Sprintf(":%d", cfg.WEB_PORT),
  }
  http.HandleFunc("/api/v1/birth-days/", router.HandleBirthDayRequest)
  // finally disconnect DB
  db, _ := rdb.DB()
  defer db.Close()
  // run Web Server
  server.ListenAndServe()
}

// create gorm instance
func newRDB(cfg *config.Config) (*gorm.DB) {
  rdb_interface := domain_rdb.NewRdbFactory(cfg.DB_RDBMS)
  return rdb_interface.ConnectDB()
}
