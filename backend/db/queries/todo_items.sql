-- name: GetTodoItem :one
SELECT id, task_id, title, completed, position, created_at, updated_at
FROM todo_items
WHERE id = $1;

-- name: CreateTodoItem :one
INSERT INTO todo_items (id, task_id, title, completed, position)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, task_id, title, completed, position, created_at, updated_at;

-- name: UpdateTodoItem :one
UPDATE todo_items
SET task_id = $2,
    title = $3,
    completed = $4,
    position = $5,
    updated_at = now()
WHERE id = $1
RETURNING id, task_id, title, completed, position, created_at, updated_at;

-- name: DeleteTodoItem :exec
DELETE FROM todo_items
WHERE id = $1;
