package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

// Create status repository
func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) FindPublicTimelines(ctx context.Context, tx *sqlx.Tx, only_media bool, since_id, limit int) (*object.Timeline, error) {
	query := `
		SELECT account.id, account.username, account.password_hash, account.display_name, account.avatar, account.header, account.note, account.create_at,
			status.id, status.account_id, status.url, status.content, status.created_at 
		FROM account 
		JOIN status ON account.id = status.account_id 
		WHERE ? <= status.account_id 
		ORDER BY status.created_at DESC 
		LIMIT ?`
	rows, err := t.db.QueryContext(ctx, query, since_id, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}
	defer rows.Close()

	var timeline object.Timeline

	for rows.Next() {
		var sta object.Status
		err := rows.Scan(&sta.Account.ID, &sta.Account.Username, &sta.Account.PasswordHash, &sta.Account.DisplayName, &sta.Account.Avatar,
			&sta.Account.Header, &sta.Account.Note, &sta.Account.CreateAt, &sta.ID, &sta.Account.ID, &sta.URL, &sta.Content, &sta.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		timeline.Status = append(timeline.Status, &sta)
	}

	return &timeline, nil
}
