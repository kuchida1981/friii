package entity

import (
	"time"

	"github.com/google/uuid"
)

type Partner struct {
	RegistrationNumber *string
	ValidTo            *time.Time
	ValidFrom          time.Time
	CreatedAt          time.Time
	Name               string
	ID                 uuid.UUID
}
