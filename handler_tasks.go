package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dominicgerman/mulberry_api/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerTasksCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title     string `json:"title"`
		Notes     string `json:"notes"`
		DueDate   string `json:"due_date"`
		Frequency string `json:"frequency"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	notes := sql.NullString{}
	if params.Notes != "" {
		notes.String = params.Notes
		notes.Valid = true
	}

	dueDate, err := time.Parse("2006-01-02", params.DueDate)
	if err != nil {
		http.Error(w, "Invalid due date format", http.StatusBadRequest)
		return
	}

	task, err := cfg.DB.CreateTask(r.Context(), database.CreateTaskParams{
		ID:          uuid.New(),
		UserID:      user.ID,
		Title:       params.Title,
		Notes:       notes,
		NextDueDate: dueDate,
		Frequency:   params.Frequency,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coudln't create task: "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseTaskToTask(task))

}

func (cfg *apiConfig) handlerTasksGet(w http.ResponseWriter, r *http.Request) {
	tasks, err := cfg.DB.GetTasks(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get tasks: "+err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, databaseTasksToTasks(tasks))
}

// func handlerTaskUpdate() {}

// func updateNextDueDate() {}

func frequencyToDuration(frequency string) (time.Duration, error) {
	switch frequency {
	case "daily":
		return 24 * time.Hour, nil
	case "weekly":
		return 7 * 24 * time.Hour, nil
	case "biweekly":
		return 14 * 24 * time.Hour, nil
	case "monthly":
		return 30 * 24 * time.Hour, nil
	case "quarterly":
		return 90 * 24 * time.Hour, nil
	case "yearly":
		return 365 * 24 * time.Hour, nil
	case "annually":
		return 365 * 24 * time.Hour, nil
	default:
		return 0, errors.New("unknown frequency")
	}
}
