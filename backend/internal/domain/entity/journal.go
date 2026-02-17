package entity

import (
	"time"

	"github.com/google/uuid"
)

type JournalEntry struct {
	OriginalEntryID *uuid.UUID
	Lines           []*JournalLine
	CreatedAt       time.Time
	TransactionDate time.Time
	Description     string
	ID              uuid.UUID
}

func (je *JournalEntry) CreateReversal(at time.Time) *JournalEntry {
	reversalID := uuid.New()
	reversal := &JournalEntry{
		ID:              reversalID,
		TransactionDate: je.TransactionDate,
		Description:     "【取消】" + je.Description,
		OriginalEntryID: &je.ID,
		CreatedAt:       at,
	}

	for _, line := range je.Lines {
		reversalLine := &JournalLine{
			ID:            uuid.New(),
			EntryID:       reversalID,
			Side:          line.OppositeSide(),
			AccountItemID: line.AccountItemID,
			Amount:        line.Amount,
			PartnerID:     line.PartnerID,
			CreatedAt:     at,
		}
		reversal.Lines = append(reversal.Lines, reversalLine)
	}

	return reversal
}

type JournalLine struct {
	PartnerID     *uuid.UUID
	CreatedAt     time.Time
	Amount        float64
	AccountItemID uuid.UUID
	EntryID       uuid.UUID
	ID            uuid.UUID
	Side          string
}

func (jl *JournalLine) OppositeSide() string {
	if jl.Side == "DEBIT" {
		return "CREDIT"
	}
	return "DEBIT"
}
