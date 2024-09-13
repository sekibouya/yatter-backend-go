package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, account object.Account, content string) (*AddStatusDTO, error)
	FindStatusByID(ctx context.Context, statusId int) (*GetStatusDTO, error)
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

func (s *status) AddStatus(ctx context.Context, account object.Account, content string) (*AddStatusDTO, error) {
	sta := object.NewStatus(account, content)

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

	statusId, err := s.statusRepo.AddStatus(ctx, tx, sta)
	if err != nil {
		return nil, err
	}
	sta.ID = *statusId
	return &AddStatusDTO{Status: sta}, nil
}

func (s *status) FindStatusByID(ctx context.Context, statusId int) (*GetStatusDTO, error) {
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
	status, err := s.statusRepo.FindStatusByID(ctx, tx, statusId)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{Status: status}, nil
}
