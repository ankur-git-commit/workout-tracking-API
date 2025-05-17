package main

import (
	"net/http"
	"time"

	"github.com/ankur-git-commit/workout-tracking-API/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("we are running our app")

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
