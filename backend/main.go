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
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db)
	movieController := controller.NewMovieController(movieService)

	// Cast Services
	castRepository := repository.NewCastRepository()
	castService := service.NewCastService(castRepository, db)
	castController := controller.NewCastController(castService)

	// Rating Services
	ratingRepository := repository.NewRatingRepository()
	ratingService := service.NewRatingService(ratingRepository, db)
	ratingController := controller.NewRatingController(ratingService)

	router := httprouter.New()
	router.GET("/api/actors", actorController.ListAll)
	router.GET("/api/casts", castController.ListAll)
	router.GET("/api/movies/:movieId", movieController.ListById)
	router.POST("/api/ratings", ratingController.SaveRating)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
