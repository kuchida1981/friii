package repository

import "github.com/kuchida1981/friii/internal/domain/entity"

//go:generate moq -out journal_mock.go . JournalRepository

type JournalRepository interface {
	Save(entry *entity.JournalEntry) error
	FindAll() ([]*entity.JournalEntry, error)
}
