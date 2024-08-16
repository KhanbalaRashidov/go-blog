package store

import (
	"context"
	"database/sql"
)

type Tag struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type TagStore struct {
	db *sql.DB
}

func (t *TagStore) Create(ctx context.Context, tag *Tag) error {
	query := `
		INSERT INTO tags ( name)
		VALUES ($1)
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := t.db.QueryRowContext(
		ctx,
		query,
		tag.Name,
	).Scan(
		&tag.Id,
		&tag.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
