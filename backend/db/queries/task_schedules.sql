-- name: GetTaskSchedule :one
SELECT id, task_id, title, description, location, start_at, end_at, due_at, created_at, updated_at
FROM task_schedules
WHERE id = $1;

-- name: CreateTaskSchedule :one
INSERT INTO task_schedules (id, task_id, title, description, location, start_at, end_at, due_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, task_id, title, description, location, start_at, end_at, due_at, created_at, updated_at;

-- name: UpdateTaskSchedule :one
UPDATE task_schedules
SET task_id = $2,
    title = $3,
    description = $4,
    location = $5,
    start_at = $6,
    end_at = $7,
    due_at = $8,
    updated_at = now()
WHERE id = $1
RETURNING id, task_id, title, description, location, start_at, end_at, due_at, created_at, updated_at;

-- name: DeleteTaskSchedule :exec
DELETE FROM task_schedules
WHERE id = $1;
