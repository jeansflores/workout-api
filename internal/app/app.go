package app

import (
	"log"
	"net/http"
	"os"

	"github.com/jeansflores/workout-api/internal/api"
)

 type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}