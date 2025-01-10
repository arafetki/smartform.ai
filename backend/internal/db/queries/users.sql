-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (id,first_name,last_name,email,avatar_url,is_email_verified,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8);

-- name: UpdateUser :execrows
UPDATE users
SET
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name),
    email = COALESCE(sqlc.narg('email'), email),
    avatar_url = COALESCE(sqlc.narg('avatar_url'), avatar_url),
    is_email_verified = COALESCE(sqlc.narg('is_email_verified'), is_email_verified)
WHERE id = sqlc.arg('id');

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1;

