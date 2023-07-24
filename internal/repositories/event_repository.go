package repositories

import (
	"fmt"

	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/dtos"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db             *gorm.DB
	userRepository *UserRepository // Add userRepository instance here
}

type UserRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	userRepo := NewUserRepository(db)

	return &EventRepository{
		db:             db,
		userRepository: userRepo,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ru *UserRepository) findOne(id uuid.UUID) (models.User, error) {
	var user models.User
	result := ru.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
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

func (r *EventRepository) CreateEvent(eventRequest models.Event) (models.Event, error) {
	result := r.db.Create(&eventRequest)
	if result.Error != nil {
		return models.Event{}, result.Error
	}
	return eventRequest, nil
}

func (r *EventRepository) AddParticipant(participantRequest dtos.ParticipantDTO) (models.Event, error) {
	user, err := r.userRepository.findOne(participantRequest.ParticipantId)
	if err != nil {
		return models.Event{}, fmt.Errorf("user with ID %d not found", participantRequest.ParticipantId)
	}

	var event models.Event
	result := r.db.Where("id = ?", participantRequest.EventId).First(&event)
	if result.Error != nil {
		return models.Event{}, result.Error
	}

	event.Participants = append(event.Participants, user)
	result = r.db.Save(&event)
	if result.Error != nil {
		return models.Event{}, result.Error
	}
	return event, nil
}
