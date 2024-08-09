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
		return fmt.Errorf("failed to insert account: %w", err)
	}

	return nil
}

func (s *status) FindStatusByID(ctx context.Context, tx *sqlx.Tx, acc_id int) ([]object.Timeline, error) {
	rows, err := s.db.QueryContext(ctx, "select * from account,status where account.id = status.account_id and status.account_id = ?", acc_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}
	defer rows.Close()

	entities := make([]object.Timeline, 0)

	for rows.Next() {
		var sta object.Status
		var acc object.Account
		err := rows.Scan(&acc.ID, &acc.Username, &acc.PasswordHash, &acc.DisplayName, &acc.Avatar, &acc.Header, &acc.Note, &acc.CreateAt,
			&sta.ID, &sta.AccountID, &sta.URL, &sta.Content, &sta.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		var entity object.Timeline
		entity.Account = &acc
		entity.Status = &sta
		entities = append(entities, entity)
	}

	return entities, nil
}
