-- name: CreatePost :one
INSERT INTO posts (title, description, content, is_draft)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = ?;

-- name: GetPostByTitle :one
SELECT * FROM posts
WHERE title = ?;

-- name: RemovePost :exec
DELETE FROM posts
WHERE id = ?;
