-- name: CreateFeedFollow :one
INSERT INTO feedFollow (feed_id, user_id) VALUES ($1, $2) RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feedFollow WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feedFollow WHERE feed_id = $1 AND user_id = $2;
