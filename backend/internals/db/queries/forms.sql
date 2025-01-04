-- name: GetFormWithSettings :one
SELECT
    f.id,
    f.user_id,
    f.title,
    f.description,
    f.fields,
    f.view_count,
    f.published,
    f.created_at,
    f.updated_at,
    sqlc.embed(s)
FROM core.forms as f
JOIN core.form_settings as s
ON f.settings_id = s.id
WHERE f.id = $1;

-- name: ListFormsForUser :many
SELECT
    id,
    user_id,
    settings_id,
    title,
    description,
    view_count,
    published,
    created_at,
    updated_at
FROM core.forms
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateForm :exec
INSERT INTO core.forms (user_id,settings_id,title,description,fields,published)
VALUES ($1,$2,$3,$4,$5,$6);

-- name: UpdateForm :execrows
UPDATE core.forms
SET
    settings_id = COALESCE(sqlc.narg('settings_id'), settings_id),
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    fields = COALESCE(sqlc.narg('fields'), fields),
    view_count = COALESCE(sqlc.narg('view_count'), view_count),
    published = COALESCE(sqlc.narg('published'), published)
WHERE id = sqlc.arg('id') AND user_id=sqlc.arg('user_id');

-- name: DeleteFormsByOwner :execrows
DELETE FROM core.forms WHERE id=ANY(sqlc.arg('ids')) AND user_id=sqlc.arg('user_id');
