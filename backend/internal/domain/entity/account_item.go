package entity

import (
	"time"

	"github.com/google/uuid"
)

type AccountCategory struct {
	Name       string
	NormalSide string
	ReportType string
	ID         uuid.UUID
}

type AccountItem struct {
	Category   *AccountCategory
	ValidTo    *time.Time
	ValidFrom  time.Time
	CreatedAt  time.Time
	Code       string
	Name       string
	ID         uuid.UUID
	CategoryID uuid.UUID
}

func (ai *AccountItem) UpdateName(newName string, at time.Time) (old, next *AccountItem) {
	ai.ValidTo = &at
	next = &AccountItem{
		ID:         ai.ID,
		CategoryID: ai.CategoryID,
		Code:       ai.Code,
		Name:       newName,
		ValidFrom:  at,
		CreatedAt:  time.Now(),
	}
	return ai, next
}
