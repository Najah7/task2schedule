-- name: GetTodoList :one
SELECT id, user_id, list_date, created_at, updated_at
FROM todo_lists
WHERE id = $1;

-- name: GetTodoListByUserAndDate :one
SELECT id, user_id, list_date, created_at, updated_at
FROM todo_lists
WHERE user_id = $1
  AND list_date = $2;

-- name: ListTodoListsByUser :many
SELECT id, user_id, list_date, created_at, updated_at
FROM todo_lists
WHERE user_id = $1
ORDER BY list_date DESC;

-- name: CreateTodoList :one
INSERT INTO todo_lists (id, user_id, list_date)
VALUES ($1, $2, $3)
RETURNING id, user_id, list_date, created_at, updated_at;

-- name: DeleteTodoList :exec
DELETE FROM todo_lists
WHERE id = $1;

-- name: ListTodoListItems :many
SELECT todo_list_id, todo_item_id, position, created_at
FROM todo_list_items
WHERE todo_list_id = $1
ORDER BY position ASC;

-- name: AddTodoListItem :one
INSERT INTO todo_list_items (todo_list_id, todo_item_id, position)
VALUES ($1, $2, $3)
RETURNING todo_list_id, todo_item_id, position, created_at;

-- name: UpdateTodoListItemPosition :one
UPDATE todo_list_items
SET position = $3
WHERE todo_list_id = $1
  AND todo_item_id = $2
RETURNING todo_list_id, todo_item_id, position, created_at;

-- name: RemoveTodoListItem :exec
DELETE FROM todo_list_items
WHERE todo_list_id = $1
  AND todo_item_id = $2;

-- name: ListTodoListTaskSchedules :many
SELECT tlts.todo_list_id, tlts.task_schedule_id, tlts.created_at
FROM todo_list_task_schedules AS tlts
JOIN task_schedules AS ts ON ts.id = tlts.task_schedule_id
WHERE tlts.todo_list_id = $1
ORDER BY ts.start_at ASC;

-- name: AddTodoListTaskSchedule :one
INSERT INTO todo_list_task_schedules (todo_list_id, task_schedule_id)
VALUES ($1, $2)
RETURNING todo_list_id, task_schedule_id, created_at;

-- name: RemoveTodoListTaskSchedule :exec
DELETE FROM todo_list_task_schedules
WHERE todo_list_id = $1
  AND task_schedule_id = $2;
