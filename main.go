package main

import (
  "fmt"
  "log"
  "net/http"
  "gorm.io/gorm"
  "birthday-reminder/config"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  domain_rdb "birthday-reminder/packages/Domain/Domain/Rdb"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  infra_repo_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)


var rdb = connectDB()
var repo_birth_day = infra_repo_birth_day.NewBirthDayRepository(rdb)
var controller_birthday = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(repo_birth_day),
  app_birth_day.NewCreateBirthDayInteractor(repo_birth_day),
)
var router = routes.NewRouter(controller_birthday)


func main() {
  cfg, err := config.NewConfig()
  if err != nil {
    log.Fatalf("failed to load config. %v", err)
  }
  server := http.Server{
    Addr: fmt.Sprintf(":%d", cfg.WEB_PORT),
  }
  http.HandleFunc("/api/v1/birth-days/", router.HandleBirthDayRequest)
  // DB切断
  db, _ := rdb.DB()
  defer db.Close()
  // WebサーバListen
  server.ListenAndServe()
}

func connectDB() (rdb *gorm.DB) {
  cfg, err := config.NewConfig()
  if err != nil {
    log.Fatalf("failed load config. %v", err)
  }
  rdb_interface := domain_rdb.NewRdbFactory(cfg.DB_RDBMS)
  rdb = rdb_interface.ConnectDB()
  return
}