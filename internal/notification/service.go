package notification

import (
	"context"

	"github.com/digitalhouse-dev/dh-kit/logger"
)

type (
	Filters struct {
		ID         []string
		UserID     []string
		NotifyType []string
	}

	Service interface {
		Create(ctx context.Context, user_id, notifyType, message string) (*Notification, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error)
	}

	service struct {
		repo   Repository
		logger logger.Logger
	}
)

func NewService(repo Repository, logger logger.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s service) Create(ctx context.Context, user_id, notifyType, message string) (*Notification, error) {
	notify := &Notification{
		UserID:     user_id,
		NotifyType: notifyType,
		Message:    message,
	}

	if err := s.repo.Create(ctx, notify); err != nil {
		return nil, s.logger.CatchError(err)
	}

	return notify, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error) {

	ns, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return ns, nil
}
