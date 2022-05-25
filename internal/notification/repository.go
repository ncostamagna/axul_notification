package notification

import (
	"context"

	"github.com/digitalhouse-dev/dh-kit/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, notification *Notification) error
	GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB, log logger.Logger) Repository {
	return &repo{db}
}

func (r repo) Create(ctx context.Context, notification *Notification) error {
	notification.ID = uuid.New().String()
	return r.db.Create(&notification).Error
}
func (r repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error) {
	var tx *gorm.DB
	var n []Notification
	tx = r.db.WithContext(ctx).Model(&n)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&n)
	if result.Error != nil {
		return nil, result.Error
	}
	return n, nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	return tx
}
