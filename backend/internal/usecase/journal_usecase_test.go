package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/kuchida1981/friii/internal/domain/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJournalRepository struct {
	mock.Mock
}

func (m *MockJournalRepository) Save(entry *entity.JournalEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

func (m *MockJournalRepository) FindByID(id uuid.UUID) (*entity.JournalEntry, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.JournalEntry), args.Error(1)
}

func (m *MockJournalRepository) List(from, to time.Time) ([]*entity.JournalEntry, error) {
	args := m.Called(from, to)
	return args.Get(0).([]*entity.JournalEntry), args.Error(1)
}

func TestJournalUsecase_CreateEntry(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	accountID := uuid.New()
	inputs := []JournalLineInput{
		{Side: "DEBIT", AccountItemID: accountID, Amount: 1000},
		{Side: "CREDIT", AccountItemID: accountID, Amount: 1000},
	}

	repo.On("Save", mock.Anything).Return(nil)

	entry, err := uc.CreateEntry(time.Now(), "Test", inputs)

	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Len(t, entry.Lines, 2)
	repo.AssertExpectations(t)
}

func TestJournalUsecase_CreateEntry_SaveError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	inputs := []JournalLineInput{
		{Side: "DEBIT", AccountItemID: uuid.New(), Amount: 1000},
		{Side: "CREDIT", AccountItemID: uuid.New(), Amount: 1000},
	}

	repo.On("Save", mock.Anything).Return(fmt.Errorf("db error"))

	entry, err := uc.CreateEntry(time.Now(), "Test", inputs)

	assert.Error(t, err)
	assert.Nil(t, entry)
	assert.Contains(t, err.Error(), "failed to save")
}

func TestJournalUsecase_CreateEntry_BalanceError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	accountID := uuid.New()
	inputs := []JournalLineInput{
		{Side: "DEBIT", AccountItemID: accountID, Amount: 1000},
		{Side: "CREDIT", AccountItemID: accountID, Amount: 2000},
	}

	entry, err := uc.CreateEntry(time.Now(), "Test", inputs)

	assert.Error(t, err)
	assert.Nil(t, entry)
	assert.Contains(t, err.Error(), "do not match")
}

func TestJournalUsecase_CreateEntry_InsufficientLines(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	inputs := []JournalLineInput{
		{Side: "DEBIT", AccountItemID: uuid.New(), Amount: 1000},
	}

	entry, err := uc.CreateEntry(time.Now(), "Test", inputs)

	assert.Error(t, err)
	assert.Nil(t, entry)
	assert.Contains(t, err.Error(), "at least 2 lines")
}

func TestJournalUsecase_CreateEntry_InvalidSide(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	inputs := []JournalLineInput{
		{Side: "INVALID", AccountItemID: uuid.New(), Amount: 1000},
		{Side: "CREDIT", AccountItemID: uuid.New(), Amount: 1000},
	}

	entry, err := uc.CreateEntry(time.Now(), "Test", inputs)

	assert.Error(t, err)
	assert.Nil(t, entry)
	assert.Contains(t, err.Error(), "invalid side")
}

func TestJournalUsecase_UpdateEntry(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)

	originalID := uuid.New()
	accountID := uuid.New()
	original := &entity.JournalEntry{
		ID:              originalID,
		TransactionDate: time.Now(),
		Description:     "Original",
		Lines: []*entity.JournalLine{
			{ID: uuid.New(), Side: "DEBIT", AccountItemID: accountID, Amount: 1000},
			{ID: uuid.New(), Side: "CREDIT", AccountItemID: accountID, Amount: 1000},
		},
	}

	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Return(nil)

	inputs := []JournalLineInput{
		{Side: "DEBIT", AccountItemID: accountID, Amount: 1200},
		{Side: "CREDIT", AccountItemID: accountID, Amount: 1200},
	}

	rev, next, err := uc.UpdateEntry(originalID, time.Now(), "Updated", inputs)

	assert.NoError(t, err)
	assert.NotNil(t, rev)
	assert.NotNil(t, next)
	assert.Equal(t, originalID, *rev.OriginalEntryID)
	assert.Equal(t, rev.ID, *next.OriginalEntryID)
	repo.AssertExpectations(t)
}

func TestJournalUsecase_UpdateEntry_FindError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	repo.On("FindByID", originalID).Return((*entity.JournalEntry)(nil), fmt.Errorf("find error"))

	_, _, err := uc.UpdateEntry(originalID, time.Now(), "Updated", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find")
}

func TestJournalUsecase_UpdateEntry_NotFound(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	repo.On("FindByID", originalID).Return((*entity.JournalEntry)(nil), nil)

	_, _, err := uc.UpdateEntry(originalID, time.Now(), "Updated", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestJournalUsecase_UpdateEntry_ReversalSaveError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	original := &entity.JournalEntry{ID: originalID}
	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Return(fmt.Errorf("save error"))

	_, _, err := uc.UpdateEntry(originalID, time.Now(), "Updated", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save reversal")
}

func TestJournalUsecase_UpdateEntry_InvalidNewEntry(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	original := &entity.JournalEntry{ID: originalID}
	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Return(nil) // Reversal save success

	// 1 line is invalid
	inputs := []JournalLineInput{{Side: "DEBIT", Amount: 1000}}

	rev, next, err := uc.UpdateEntry(originalID, time.Now(), "Updated", inputs)

	assert.Error(t, err)
	assert.NotNil(t, rev)
	assert.Nil(t, next)
	assert.Contains(t, err.Error(), "at least 2 lines")
}

func TestJournalUsecase_UpdateEntry_NewEntryBalanceError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	original := &entity.JournalEntry{ID: originalID}
	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Return(nil) // Reversal save success

	inputs := []JournalLineInput{
		{Side: "DEBIT", Amount: 1000},
		{Side: "CREDIT", Amount: 2000},
	}

	rev, next, err := uc.UpdateEntry(originalID, time.Now(), "Updated", inputs)

	assert.Error(t, err)
	assert.NotNil(t, rev)
	assert.Nil(t, next)
	assert.Contains(t, err.Error(), "do not match")
}

func TestJournalUsecase_UpdateEntry_NewEntrySaveError(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	original := &entity.JournalEntry{ID: originalID}
	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Once().Return(nil)                    // Reversal success
	repo.On("Save", mock.Anything).Once().Return(fmt.Errorf("db error")) // New entry fail

	inputs := []JournalLineInput{
		{Side: "DEBIT", Amount: 1000},
		{Side: "CREDIT", Amount: 1000},
	}

	rev, next, err := uc.UpdateEntry(originalID, time.Now(), "Updated", inputs)

	assert.Error(t, err)
	assert.NotNil(t, rev)
	assert.Nil(t, next)
	assert.Contains(t, err.Error(), "failed to save new entry")
}

func TestJournalUsecase_UpdateEntry_InvalidInputSide(t *testing.T) {
	repo := new(MockJournalRepository)
	uc := NewJournalUsecase(repo)
	originalID := uuid.New()
	original := &entity.JournalEntry{ID: originalID}
	repo.On("FindByID", originalID).Return(original, nil)
	repo.On("Save", mock.Anything).Once().Return(nil) // Reversal success

	inputs := []JournalLineInput{
		{Side: "INVALID", Amount: 1000},
		{Side: "CREDIT", Amount: 1000},
	}

	rev, _, err := uc.UpdateEntry(originalID, time.Now(), "Updated", inputs)

	assert.Error(t, err)
	assert.NotNil(t, rev)
	assert.Contains(t, err.Error(), "invalid side")
}
