-- name: GetForm :one
SELECT *
FROM core.forms
WHERE id = $1;

-- name: ListFormsForUser :many
SELECT
    id,
    user_id,
    title,
    description,
    published,
    created_at,
    updated_at
FROM core.forms
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateForm :exec
INSERT INTO core.forms (user_id,title,description,fields,published)
VALUES ($1,$2,$3,$4,$5);

-- name: UpdateForm :execrows
UPDATE core.forms
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    fields = COALESCE(sqlc.narg('fields'), fields),
    published = COALESCE(sqlc.narg('published'), published)
WHERE id = sqlc.arg('id');

-- name: DeleteForm :exec
DELETE FROM core.forms WHERE id=$1;
