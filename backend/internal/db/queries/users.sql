-- name: GetUser :one
SELECT * FROM core.users
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO core.users (id,avatar_url,is_verified,created_at) VALUES ($1,$2,$3,$4);

-- name: UpdateUser :execrows
UPDATE core.users
SET
    avatar_url = COALESCE(sqlc.narg('avatar_url'), avatar_url),
    is_verified = COALESCE(sqlc.narg('is_verified'), is_verified)
WHERE id = sqlc.arg('id');

-- name: DeleteUser :execrows
DELETE FROM core.users WHERE id = $1;

