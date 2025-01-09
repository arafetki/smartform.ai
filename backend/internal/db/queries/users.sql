-- name: GetUser :one
SELECT * FROM core.users
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO core.users (id,is_verified,created_at) VALUES ($1,$2,$3);

-- name: UpdateUser :execrows
UPDATE core.users
SET
    is_verified = COALESCE(sqlc.narg('is_verified'), is_verified)
WHERE id = sqlc.arg('id');

-- name: DeleteUser :execrows
DELETE FROM core.users WHERE id = $1;

