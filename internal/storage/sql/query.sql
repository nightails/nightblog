-- name: CreatePost :one
INSERT INTO posts (title, description, content, is_draft)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = ?;
