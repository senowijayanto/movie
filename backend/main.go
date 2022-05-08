package main

import (
	"backend/app"
	"backend/controller"
	"backend/helper"
	"backend/repository"
	"backend/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	actorRepository := repository.NewActorRepository()
	actorService := service.NewActorService(actorRepository, db)
	actorController := controller.NewActorController(actorService)
	router := app.NewRouter(actorController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
