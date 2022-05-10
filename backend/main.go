package main

import (
	"backend/app"
	"backend/controller"
	"backend/helper"
	"backend/repository"
	"backend/service"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Movie Backend!")
	db := app.NewDB()

	movieHandler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintf(w, "Movie Backend")
	}
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
	router.GET("/", movieHandler)
	router.GET("/api/actors", actorController.ListAll)
	router.GET("/api/casts", castController.ListAll)
	router.GET("/api/movies/:movieId", movieController.ListById)
	router.GET("/api/movies", movieController.ListAll)
	router.POST("/api/ratings", ratingController.SaveRating)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
