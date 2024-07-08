package main

import (
	"database/sql"
	"time"

	"github.com/dominicgerman/mulberry_api/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Notes       *string   `json:"notes"`
	Frequency   string    `json:"frequency"`
	NextDueDate time.Time `json:"next_due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databaseTaskToTask(task database.Task) Task {
	return Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Notes:       nullStringToStringPtr(task.Notes),
		Frequency:   task.Frequency,
		NextDueDate: task.NextDueDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func databaseTasksToTasks(tasks []database.Task) []Task {
	result := make([]Task, len(tasks))
	for i, task := range tasks {
		result[i] = databaseTaskToTask(task)
	}
	return result
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
