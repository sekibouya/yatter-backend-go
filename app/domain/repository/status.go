package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) error
}