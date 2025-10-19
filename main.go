package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/jeansflores/workout-api/internal/app"
	"github.com/jeansflores/workout-api/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the HTTP server on")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application started")

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Starting server on port %d\n", port)


	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatalf("Could not start server: %s\n", err.Error())
	}
}


