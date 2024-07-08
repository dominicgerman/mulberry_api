-- name: CreateTask :one
INSERT INTO tasks (id, user_id, title, notes, frequency, next_due_date, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetTasks :many
SELECT * from tasks;