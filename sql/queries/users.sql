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

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: GetAllUsers :many
SELECT *
FROM users;