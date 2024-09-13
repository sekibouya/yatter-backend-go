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

func (s *status) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) (*int, error) {
	_, err := tx.ExecContext(ctx, "insert into status (account_id, url, content, created_at) values (?, ?, ?, ?)",
		status.Account.ID, status.URL, status.Content, status.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert status: %w", err)
	}
	var statusId int
	tx.QueryRowContext(ctx, "select id from status order by id desc limit 1;").Scan(&statusId)
	return &statusId, nil
}

func (s *status) FindStatusByID(ctx context.Context, tx *sqlx.Tx, statusId int) (*object.Status, error) {
	row := s.db.QueryRowxContext(ctx, "select * from account,status where account.id = status.account_id and status.id = ?", statusId)
	status := new(object.Status)
	err := row.Scan(&status.Account.ID, &status.Account.Username, &status.Account.PasswordHash, &status.Account.DisplayName, &status.Account.Avatar,
		&status.Account.Header, &status.Account.Note, &status.Account.CreateAt, &status.ID, &status.Account.ID, &status.URL, &status.Content, &status.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return status, nil
}
