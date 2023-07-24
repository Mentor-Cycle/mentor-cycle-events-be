package services

import (
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/dtos"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/models"
	"github.com/Mentor-Cycle/mentor-cycle-events-be/internal/repositories"
	"github.com/google/uuid"
)

type EventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository repositories.EventRepository) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}

func (s *EventService) FindAllEvents() ([]models.Event, error) {
	return s.eventRepository.FindAll()
}

func (s *EventService) CreateEvent(eventRequest dtos.EventDTO) (models.Event, error) {

	eventRequestParsed := models.Event{
		ID:          uuid.New().String(),
		Title:       eventRequest.Title,
		Description: eventRequest.Description,
		MentorID:    eventRequest.MentorID,
		StartDate:   eventRequest.StartDate,
		EndDate:     eventRequest.EndDate,
	}

	return s.eventRepository.CreateEvent(eventRequestParsed)
}

func (s *EventService) AddParticipant(participantRequest dtos.ParticipantDTO) (models.Event, error) {
	return s.eventRepository.AddParticipant(participantRequest)
}
