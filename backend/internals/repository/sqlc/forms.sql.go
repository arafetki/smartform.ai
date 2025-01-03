// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: forms.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFormWithSettings = `-- name: GetFormWithSettings :one
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
    s.id, s.background_color, s.foreground_color, s.primary_color, s.created_at
FROM core.forms as f
JOIN core.form_settings as s
ON f.settings_id = s.id
WHERE f.id = $1
`

type GetFormWithSettingsRow struct {
	ID           pgtype.UUID        `json:"id"`
	UserID       pgtype.UUID        `json:"user_id"`
	Title        string             `json:"title"`
	Description  pgtype.Text        `json:"description"`
	Fields       []byte             `json:"fields"`
	ViewCount    int64              `json:"view_count"`
	Published    bool               `json:"published"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	FormSettings FormSettings       `json:"form_settings"`
}

func (q *Queries) GetFormWithSettings(ctx context.Context, id pgtype.UUID) (GetFormWithSettingsRow, error) {
	row := q.db.QueryRow(ctx, getFormWithSettings, id)
	var i GetFormWithSettingsRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Fields,
		&i.ViewCount,
		&i.Published,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FormSettings.ID,
		&i.FormSettings.BackgroundColor,
		&i.FormSettings.ForegroundColor,
		&i.FormSettings.PrimaryColor,
		&i.FormSettings.CreatedAt,
	)
	return i, err
}

const listFormsForUser = `-- name: ListFormsForUser :many
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
ORDER BY created_at DESC
`

type ListFormsForUserRow struct {
	ID          pgtype.UUID        `json:"id"`
	UserID      pgtype.UUID        `json:"user_id"`
	SettingsID  pgtype.Int2        `json:"settings_id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	ViewCount   int64              `json:"view_count"`
	Published   bool               `json:"published"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) ListFormsForUser(ctx context.Context, userID pgtype.UUID) ([]ListFormsForUserRow, error) {
	rows, err := q.db.Query(ctx, listFormsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListFormsForUserRow{}
	for rows.Next() {
		var i ListFormsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SettingsID,
			&i.Title,
			&i.Description,
			&i.ViewCount,
			&i.Published,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}