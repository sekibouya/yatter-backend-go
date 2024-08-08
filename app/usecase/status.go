package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, acc_id int, content string) (*AddStatusDTO, error)
	FindStatusByID(ctx context.Context, acc_id int) (*GetStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type AddStatusDTO struct {
	Status *object.Status
}

type GetStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) AddStatus(ctx context.Context, acc_id int, content string) (*AddStatusDTO, error) {
	sta := object.NewStatus(acc_id, content)

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.AddStatus(ctx, tx, sta); err != nil {
		return nil, err
	}

	return &AddStatusDTO{
		Status: sta,
	}, nil
}

func (s *status) FindStatusByID(ctx context.Context, acc_id int) (*GetStatusDTO, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()
	sta, err := s.statusRepo.FindStatusByID(ctx, tx, acc_id)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: sta,
	}, nil
}
