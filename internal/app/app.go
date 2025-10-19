package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jeansflores/workout-api/internal/api"
	"github.com/jeansflores/workout-api/internal/store"
	"github.com/jeansflores/workout-api/migrations"
)

 type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB			*sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
		DB: pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}