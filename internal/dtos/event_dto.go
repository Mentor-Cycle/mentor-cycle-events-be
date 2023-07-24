package dtos

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type EventDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	MentorID    string `json:"mentor_id" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}

func ValidateEventDTO(eventDTO *EventDTO) error {
	return validate.Struct(eventDTO)
}
