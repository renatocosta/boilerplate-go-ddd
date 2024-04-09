-- name: GetLogFile :one
SELECT * FROM log_file
WHERE id = ? LIMIT 1;

-- name: ListLogFiles :many
SELECT * FROM log_file
ORDER BY path;

-- name: CreateLogFile :execresult
INSERT INTO log_file (
  id, path
) VALUES (
  ?, ?
);

-- name: DeleteLogFile :exec
DELETE FROM log_file
WHERE id = ?;