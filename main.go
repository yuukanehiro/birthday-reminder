package main

import (
  "fmt"
  "log"
  "net/http"
  "birthday-reminder/config"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  infra_repo_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)

var birth_day_repo = infra_repo_birth_day.NewBirthDayRepository()
var controller_birthday = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(birth_day_repo),
  app_birth_day.NewCreateBirthDayInteractor(birth_day_repo),
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
  server.ListenAndServe()
}
