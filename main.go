package main

import (
  //"birthday-reminder/routes"
  "net/http"
  "birthday-reminder/routes"
  "birthday-reminder/controllers"
  app_birth_day "birthday-reminder/packages/Domain/Application/BirthDay"
)

var controller_birthday = controllers.NewBirthDayController(
  app_birth_day.NewListBirthDayInteractor(),
  app_birth_day.NewCreateBirthDayInteractor(),
)
var router = routes.NewRouter(controller_birthday)

func main() {
  server := http.Server{
    Addr: ":8080",
  }
  http.HandleFunc("/birth-days/", router.HandleBirthDayRequest)
  server.ListenAndServe()
}
