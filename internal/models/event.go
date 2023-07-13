package models

import (
	"time"
)

type Event struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	MentorID     string    `gorm:"column:mentor_id" json:"mentorId"`
	Participants []User    `gorm:"many2many:EventsOnUsers;" json:"participants"`
	StartDate    string    `gorm:"column:start_date" json:"startDate"`
	MeetingLink  string    `gorm:"column:meeting_link" json:"meetingLink"`
	EndDate      string    `gorm:"column:end_date" json:"endDate"`
	Active       bool      `gorm:"default:false" json:"active"`
	Status       string    `gorm:"default:'CONFIRMED'" json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (Event) TableName() string {
	return "events"
}
