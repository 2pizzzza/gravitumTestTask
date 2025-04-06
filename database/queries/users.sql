-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (username, email)
VALUES ($1, $2)
    RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    email = $3,
    updated_at = NOW()
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1;