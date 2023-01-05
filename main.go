package main

import (
  "net/http"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
  infra_repo_birth_day "birthday-reminder/packages/Infrastructure/Repositories/BirthDay"
)


var interface_birth_day = infra_repo_birth_day.NewBirthDayRepository()
var controller_birthday = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(interface_birth_day),
  app_birth_day.NewCreateBirthDayInteractor(interface_birth_day),
)
var router = routes.NewRouter(controller_birthday)


func main() {
  server := http.Server{
    Addr: ":8080",
  }
  http.HandleFunc("/api/v1/birth-days/", router.HandleBirthDayRequest)
  server.ListenAndServe()
}
