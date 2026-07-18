-- name: GetTask :one
SELECT id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
       priority, status, created_at, updated_at
FROM tasks
WHERE id = $1;

-- name: GetTaskByTag :many
SELECT t.id, t.user_id, t.project_id, t.title, t.description, t.estimated_minutes, t.actual_minutes, t.progress,
       t.priority, t.status, t.created_at, t.updated_at
FROM tasks AS t
JOIN task_tag_assignments AS tta ON tta.task_id = t.id
WHERE tta.tag_id = $1;

-- name: GetTaskByStatus :many
SELECT id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
       priority, status, created_at, updated_at
FROM tasks
WHERE status = $1;

-- name: GetTaskByPriority :many
SELECT id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
       priority, status, created_at, updated_at
FROM tasks
WHERE priority = $1;

-- name: GetTaskByProject :many
SELECT id, user_id, project_id, title, description, estimated_minutes, actual_minutes, progress,
       priority, status, created_at, updated_at
FROM tasks
WHERE project_id = sqlc.arg(project_id)::text;

-- name: GetTaskByProjectType :many
SELECT t.id, t.user_id, t.project_id, t.title, t.description, t.estimated_minutes, t.actual_minutes, t.progress,
       t.priority, t.status, t.created_at, t.updated_at
FROM tasks AS t
JOIN projects AS p ON p.id = t.project_id
WHERE p.type = $1;

-- name: GetTaskByFrequency :many
SELECT t.id, t.user_id, t.project_id, t.title, t.description, t.estimated_minutes, t.actual_minutes, t.progress,
       t.priority, t.status, t.created_at, t.updated_at
FROM tasks AS t
WHERE EXISTS (
    SELECT 1
    FROM todo_items AS ti
    JOIN todo_item_frequencies AS tif ON tif.todo_item_id = ti.id
    WHERE ti.task_id = t.id
      AND tif.frequency = sqlc.arg(frequency)::text
)
OR EXISTS (
    SELECT 1
    FROM task_schedules AS ts
    JOIN task_schedule_frequencies AS tsf ON tsf.task_schedule_id = ts.id
    WHERE ts.task_id = t.id
      AND tsf.frequency = sqlc.arg(frequency)::text
);
