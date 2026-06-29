-- name: CreatePost :one
INSERT INTO posts (id, title, description, content, is_draft)
VALUES (?, ?, ?, ?, ?)
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

-- name: UpdatePostContent :exec
UPDATE posts
SET content = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;
