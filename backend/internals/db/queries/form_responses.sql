-- name: ListResponsesForForm :many
SELECT *
FROM core.form_responses
WHERE form_id = $1
ORDER BY created_at DESC;

-- name: CreateResponsesForForm :exec
INSERT INTO core.form_responses (form_id, data)
VALUES ($1,$2);