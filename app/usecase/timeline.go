package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

const maxLimit = 80

type Timeline interface {
	FindPublicTimelines(ctx context.Context, only_media bool, since_id, limit int) (*GetTimelineDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	timelineRepo repository.Timeline
}

type GetTimelineDTO struct {
	Timelines []object.Timeline
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) FindPublicTimelines(ctx context.Context, only_media bool, since_id, limit int) (*GetTimelineDTO, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if maxLimit < limit {
		limit = maxLimit
	}
	timelines, err := t.timelineRepo.FindPublicTimelines(ctx, tx, only_media, since_id, limit)
	if err != nil {
		return nil, err
	}

	return &GetTimelineDTO{
		Timelines: timelines,
	}, nil
}
