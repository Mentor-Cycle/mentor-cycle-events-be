package models

import (
	"time"

	"gorm.io/gorm"
)

type EventsOnUsers struct {
	gorm.Model
	Event      Event     `gorm:"foreignKey:EventID"`
	EventID    string    `gorm:"primaryKey"`
	User       User      `gorm:"foreignKey:UserID"`
	UserID     string    `gorm:"primaryKey"`
	AssignedAt time.Time `gorm:"default:current_timestamp"`
	AssignedBy *string   `gorm:"column:assigned_by default:'SYSTEM'"`
}

func (EventsOnUsers) TableName() string {
	return "EventsOnUsers"
}
