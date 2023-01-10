package main

import (
  "fmt"
  "net/http"
  "gorm.io/gorm"
  "birthday-reminder/config"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  // BirthDay
  infra_repo_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
  infra_repo_Rdb "birthday-reminder/packages/Infrastructure/Repositories/Rdb"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  // Auth
  infra_repo_auth "birthday-reminder/packages/Infrastructure/Repositories/Auth"
  app_auth "birthday-reminder/packages/Domain/Application/Auth"
)

var cfg = config.NewConfig()
var rdb = newRDB(cfg)
// BirthDay
var repo_birth_day = infra_repo_birth_day.NewBirthDayRepository(rdb)
var controller_birth_day = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(repo_birth_day),
  app_birth_day.NewCreateBirthDayInteractor(repo_birth_day),
)
// Auth
var repo_auth_register = infra_repo_auth.NewAuthRegisterRepository(rdb)
var controller_auth_register = controllers.NewAuthRegisterController(
  app_auth.NewAuthRegisterInteractor(repo_auth_register),
)

var router = routes.NewRouter(
  controller_birth_day,
  controller_auth_register,
)


func main() {
  server := http.Server{
    Addr: fmt.Sprintf(":%d", cfg.WEB_PORT),
  }
  http.HandleFunc("/api/v1/user/register/", router.HandleAuthRegisterRequest)
  http.HandleFunc("/api/v1/birth-days/", router.HandleBirthDayRequest)

  // finally disconnect DB
  db, _ := rdb.DB()
  defer db.Close()
  // run Web Server
  server.ListenAndServe()
}

// create gorm instance
func newRDB(cfg *config.Config) (*gorm.DB) {
  rdb := infra_repo_Rdb.NewRdb(cfg)
  return rdb.ConnectDB()
}

