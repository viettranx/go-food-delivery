package common

import "time"

type SQLModel struct {
	ID        int        `json:"id" gorm:"primaryKey;column:id"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
