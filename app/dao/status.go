package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create status repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (a *status) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) error {
	_, err := a.db.ExecContext(ctx, "insert into status (account_id, url, content, created_at) values (?, ?, ?, ?)",
		status.AccountID, status.URL, status.Content, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("change! failed to insert account: %w", err)
	}

	return nil
}
