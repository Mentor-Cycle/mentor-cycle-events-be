package repositories

import (
	"fmt"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) FindAll() ([]models.Event, error) {
	var events []models.Event
	result := r.db.Preload("Participants").Find(&events)
	fmt.Println("result", result)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}
