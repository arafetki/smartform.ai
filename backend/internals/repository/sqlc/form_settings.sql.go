// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: form_settings.sql

package sqlc

import (
	"context"
)

const listSettings = `-- name: ListSettings :many
SELECT
    id, background_color, foreground_color, primary_color, created_at
FROM
    core.form_settings
`

func (q *Queries) ListSettings(ctx context.Context) ([]FormSettings, error) {
	rows, err := q.db.Query(ctx, listSettings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FormSettings{}
	for rows.Next() {
		var i FormSettings
		if err := rows.Scan(
			&i.ID,
			&i.BackgroundColor,
			&i.ForegroundColor,
			&i.PrimaryColor,
			&i.CreatedAt,
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
