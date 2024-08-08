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

func (s *status) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) error {
	_, err := s.db.ExecContext(ctx, "insert into status (account_id, url, content, created_at) values (?, ?, ?, ?)",
		status.AccountID, status.URL, status.Content, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("change! failed to insert account: %w", err)
	}

	return nil
}

func (s *status) FindStatusByID(ctx context.Context, tx *sqlx.Tx, acc_id int) ([]*object.Status, error) {
	rows, err := s.db.QueryContext(ctx, "select * from status where account_id = ?", acc_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}
	defer rows.Close()

	entities := make([]*object.Status, 0)

	for rows.Next() {
		var entity object.Status
		rows.Scan(&entity.ID, &entity.AccountID, &entity.URL, &entity.Content, &entity.CreatedAt)
		entities = append(entities, &entity)
	}

	return entities, nil
}
