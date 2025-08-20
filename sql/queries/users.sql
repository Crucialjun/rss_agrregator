-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, name, api_key)
VALUES (gen_random_uuid(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, $2, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE api_key = $1;

-- name: UpdateUser :one
UPDATE users SET email = $1, name = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1
RETURNING *;
