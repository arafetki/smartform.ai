-- name: GetUser :one
SELECT * FROM core.users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM core.users
ORDER BY created_at DESC;

-- name: CreateUser :exec
INSERT INTO core.users (id,email,name,phone_number,is_verified,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7);

-- name: UpdateUser :exec
UPDATE core.users
SET
    email = COALESCE(sqlc.narg('email'), email),
    name = COALESCE(sqlc.narg('name'), name),
    phone_number = COALESCE(sqlc.narg('phone_number'), phone_number),
    is_verified = COALESCE(sqlc.narg('is_verified'), is_verified),
    avatar_url = COALESCE(sqlc.narg('avatar_url'), avatar_url)
WHERE id = sqlc.arg('id');

-- name: DeleteUser :exec
DELETE FROM core.users WHERE id = $1;

-- name: CountUsers :one
SELECT count(*) FROM core.users;
