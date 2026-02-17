package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJournalEntry_CreateReversal(t *testing.T) {
	now := time.Now()
	entryID := uuid.New()
	accountID := uuid.New()

	entry := &JournalEntry{
		ID:              entryID,
		TransactionDate: now,
		Description:     "Test Entry",
		Lines: []*JournalLine{
			{
				ID:            uuid.New(),
				Side:          "DEBIT",
				AccountItemID: accountID,
				Amount:        1000,
			},
			{
				ID:            uuid.New(),
				Side:          "CREDIT",
				AccountItemID: accountID,
				Amount:        1000,
			},
		},
	}

	reversal := entry.CreateReversal(now)

	assert.NotEqual(t, entry.ID, reversal.ID)
	assert.Equal(t, entry.ID, *reversal.OriginalEntryID)
	assert.Contains(t, reversal.Description, "【取消】")
	assert.Len(t, reversal.Lines, 2)

	assert.Equal(t, "CREDIT", reversal.Lines[0].Side)
	assert.Equal(t, "DEBIT", reversal.Lines[1].Side)
	assert.Equal(t, 1000.0, reversal.Lines[0].Amount)
}
