package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountItem_UpdateName(t *testing.T) {
	now := time.Now()
	ai := &AccountItem{
		ID:        uuid.New(),
		Name:      "Old Name",
		ValidFrom: now.Add(-24 * time.Hour),
	}

	updateTime := now
	oldAI, nextAI := ai.UpdateName("New Name", updateTime)

	assert.Equal(t, ai.ID, nextAI.ID)
	assert.Equal(t, "New Name", nextAI.Name)
	assert.Equal(t, updateTime, nextAI.ValidFrom)
	assert.Equal(t, &updateTime, oldAI.ValidTo)
}
