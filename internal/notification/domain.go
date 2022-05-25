package notification

import (
	"time"

	"gorm.io/gorm"
)

// Notification entity
type Notification struct {
	ID         string         `json:"id" gorm:"not null;primary_key;unique_index"`
	UserID     string         `json:"-" gorm:"type:varchar(50);not null"`
	NotifyType string         `json:"notify_type" gorm:"type:varchar(20);not null"`
	Message    string         `json:"message" gorm:"type:varchar(200);not null"`
	Checked    bool           `json:"checked"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	CreatedAt  time.Time      `json:"-" gorm:"type:timestamp;column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"-" gorm:"type:timestamp;column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (d Notification) TableName() string {
	return "notification"
}
