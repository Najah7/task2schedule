
-- name: CreateTask :one
INSERT INTO tasks (
    id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
    priority, status
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
          priority, status, created_at, updated_at;

-- name: UpdateTask :one
UPDATE tasks
SET user_id = $2,
    project_id = $3,
    title = $4,
    description = $5,
    estimated_minutes = $6,
    actual_minutes = $7,
    progress = $8,
    priority = $9,
    status = $10,
    updated_at = now()
WHERE id = $1
RETURNING id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
          priority, status, created_at, updated_at;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
