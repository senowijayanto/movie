package main

import (
	"backend/app"
	"backend/controller"
	"backend/helper"
	"backend/repository"
	"backend/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()

	// Actor Services
	actorRepository := repository.NewActorRepository()
	actorService := service.NewActorService(actorRepository, db)
	actorController := controller.NewActorController(actorService)

	// Movie Services

	// Cast Services
	castRepository := repository.NewCastRepository()
	castService := service.NewCastService(castRepository, db)
	castController := controller.NewCastController(castService)

	router := httprouter.New()
	router.GET("/api/actors", actorController.ListAll)
	router.GET("/api/casts", castController.ListAll)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
