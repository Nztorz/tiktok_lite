-- name: CreateUser :one
INSERT INTO users(email, username, hashed_password, bio)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;