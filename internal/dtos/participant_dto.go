package dtos

import "github.com/google/uuid"

type ParticipantDTO struct {
	ParticipantId uuid.UUID `json:"participantId" binding:"required"`
	EventId       uuid.UUID `-`
}
