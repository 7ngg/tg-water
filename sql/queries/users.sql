-- name: ListUsers :many
SELECT * FROM user;

-- name: GetByTgID :one
SELECT * FROM user
WHERE user.telegram_id=?;

-- name: SaveUser :one
INSERT INTO user(telegram_id, chat_id)
VALUES(?, ?)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM user
WHERE telegram_id = ?;
