-- name: CreateUser :one
INSERT INTO users (id, username, hashed_password, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByID :one
SELECT * 
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: UpdatePasswordById :exec
UPDATE users
SET 
    hashed_password = $1,
    updated_at = NOW()
WHERE id = $2;