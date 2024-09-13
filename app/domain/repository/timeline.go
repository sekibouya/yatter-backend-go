package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	FindPublicTimelines(ctx context.Context, tx *sqlx.Tx, only_media bool, since_id, limit int) (*object.Timeline, error)
}
