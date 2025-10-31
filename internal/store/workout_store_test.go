package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestCreateWork(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:        "Morning Routine",
				Description: "A quick morning workout",
				DurationMinutes: 60,
				CaloriesBurned: 200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Push-ups",
						Sets: 3,
						Reps: IntPtr(15),
						Weight: FloatPtr(135.5), 
						Notes: "Keep back straight",
						OrderIndex: 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "workout with invalid entries",
			workout: &Workout{
				Title:        "Invalid Workout",
				Description:  "This workout has invalid entries",
				DurationMinutes: 45,
				CaloriesBurned: 150,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Invalid Exercise",
						Sets: 0,
						Reps: IntPtr(-5),
						Weight: FloatPtr(-10.0),
						Notes: "Invalid entry",
						OrderIndex: 1,
					},
					{
						ExerciseName: "",
						Sets: 3,
						Reps: IntPtr(10),
						DurationSeconds: IntPtr(60),
						Weight: FloatPtr(50.0),
						Notes: "Missing exercise name",
						OrderIndex: 2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createdWorkout.Title)
			assert.Equal(t, tt.workout.Description, createdWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)
			assert.Equal(t, createdWorkout.ID, retrieved.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrieved.Entries))

			for i, entry := range retrieved.Entries {
				expectedEntry := tt.workout.Entries[i]
				assert.Equal(t, expectedEntry.ExerciseName, entry.ExerciseName)
				assert.Equal(t, expectedEntry.Sets, entry.Sets)
				assert.Equal(t, expectedEntry.OrderIndex, entry.OrderIndex)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPtr(f float64) *float64 {
	return &f
}