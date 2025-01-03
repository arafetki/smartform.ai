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