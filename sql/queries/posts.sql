-- name: CreatePost :one
INSERT INTO posts (feed_id, title, link, published_at, created_at, updated_at, url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


