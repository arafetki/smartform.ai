-- name: GetForm :one
SELECT
    id,
    user_id,
    title,
    description,
    fields,
    is_published,
    created_at,
    updated_at
FROM forms
WHERE id = $1;

-- name: GetAllFormsForUser :many
SELECT
    id,
    user_id,
    title,
    description,
    fields,
    is_published,
    created_at,
    updated_at
FROM forms
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateForm :exec
INSERT INTO forms (user_id,title,description,fields,is_published)
VALUES ($1,$2,$3,$4,$5);

-- name: UpdateForm :execrows
UPDATE forms
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    fields = COALESCE(sqlc.narg('fields'), fields),
    is_published = COALESCE(sqlc.narg('is_published'), is_published)
WHERE id = sqlc.arg('id');

-- name: DeleteForm :execrows
DELETE FROM forms WHERE id=$1;
