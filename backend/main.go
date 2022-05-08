package main

import (
	"backend/helper"
	"net/http"
)

func main() {
	// db := app.NewDB()

	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
