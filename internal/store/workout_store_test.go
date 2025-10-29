package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=workout_user password=secure_password dbname=workout_test_db port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("db: open %v", err)
	}

	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("db: migrate %v", err)
	}

	_, err = db.Exec("TRUNCATE workouts, workout_entries CASCADE")
	if err != nil {
		t.Fatalf("db: truncate %v", err)
	}

	return db
}