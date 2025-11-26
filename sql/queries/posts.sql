-- name: CreatePost :one
INSERT INTO posts (id, body, user_id, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetPostByID :one
SELECT *
FROM posts
WHERE id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts
WHERE id = $1;

-- name: EditPostBodyByID :exec
UPDATE posts
SET 
    body = $1,
    updated_at = NOW()
WHERE user_id = $2;

