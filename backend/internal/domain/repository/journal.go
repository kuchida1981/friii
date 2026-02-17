package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/kuchida1981/friii/internal/domain/entity"
)

//go:generate moq -out journal_mock.go . JournalRepository

type JournalRepository interface {
	Save(entry *entity.JournalEntry) error
	FindByID(id uuid.UUID) (*entity.JournalEntry, error)
	List(from, to time.Time) ([]*entity.JournalEntry, error)
}
