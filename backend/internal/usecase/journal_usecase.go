package usecase

import (
	"fmt"
	"time"

	"github.com/kuchida1981/friii/internal/domain/entity"
	"github.com/kuchida1981/friii/internal/domain/repository"

	"github.com/google/uuid"
)

type JournalUsecase interface {
	CreateEntry(date time.Time, description string, lines []JournalLineInput) (*entity.JournalEntry, error)
	UpdateEntry(originalID uuid.UUID, date time.Time, description string, lines []JournalLineInput) (reversal, next *entity.JournalEntry, err error)
}

type JournalLineInput struct {
	PartnerID     *uuid.UUID
	Amount        float64
	AccountItemID uuid.UUID
	Side          string
}

type journalUsecase struct {
	journalRepo repository.JournalRepository
}

func NewJournalUsecase(journalRepo repository.JournalRepository) JournalUsecase {
	return &journalUsecase{journalRepo: journalRepo}
}

func (u *journalUsecase) CreateEntry(date time.Time, description string, inputs []JournalLineInput) (*entity.JournalEntry, error) {
	if len(inputs) < 2 {
		return nil, fmt.Errorf("journal entry must have at least 2 lines")
	}

	var debitTotal, creditTotal float64
	entryID := uuid.New()
	entry := &entity.JournalEntry{
		ID:              entryID,
		TransactionDate: date,
		Description:     description,
		CreatedAt:       time.Now(),
	}

	for _, input := range inputs {
		switch input.Side {
		case "DEBIT":
			debitTotal += input.Amount
		case "CREDIT":
			creditTotal += input.Amount
		default:
			return nil, fmt.Errorf("invalid side: %s", input.Side)
		}

		line := &entity.JournalLine{
			ID:            uuid.New(),
			EntryID:       entryID,
			Side:          input.Side,
			AccountItemID: input.AccountItemID,
			Amount:        input.Amount,
			PartnerID:     input.PartnerID,
			CreatedAt:     entry.CreatedAt,
		}
		entry.Lines = append(entry.Lines, line)
	}

	// 貸借一致チェック
	if debitTotal != creditTotal {
		return nil, fmt.Errorf("debit and credit totals do not match: %f != %f", debitTotal, creditTotal)
	}

	if err := u.journalRepo.Save(entry); err != nil {
		return nil, fmt.Errorf("failed to save journal entry: %w", err)
	}

	return entry, nil
}

func (u *journalUsecase) UpdateEntry(originalID uuid.UUID, date time.Time, description string, inputs []JournalLineInput) (reversal, next *entity.JournalEntry, err error) {
	// 1. 元の仕訳を取得
	original, err := u.journalRepo.FindByID(originalID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find original entry: %w", err)
	}
	if original == nil {
		return nil, nil, fmt.Errorf("original entry not found: %s", originalID)
	}

	now := time.Now()

	// 2. 逆仕訳（赤）を作成して保存
	reversal = original.CreateReversal(now)
	if err := u.journalRepo.Save(reversal); err != nil {
		return nil, nil, fmt.Errorf("failed to save reversal entry: %w", err)
	}

	// 3. 正しい新規仕訳（黒）を作成して保存
	if len(inputs) < 2 {
		return reversal, nil, fmt.Errorf("journal entry must have at least 2 lines")
	}

	var debitTotal, creditTotal float64
	newEntryID := uuid.New()
	next = &entity.JournalEntry{
		ID:              newEntryID,
		TransactionDate: date,
		Description:     description,
		OriginalEntryID: &reversal.ID,
		CreatedAt:       now,
	}

	for _, input := range inputs {
		switch input.Side {
		case "DEBIT":
			debitTotal += input.Amount
		case "CREDIT":
			creditTotal += input.Amount
		default:
			return reversal, nil, fmt.Errorf("invalid side: %s", input.Side)
		}

		line := &entity.JournalLine{
			ID:            uuid.New(),
			EntryID:       newEntryID,
			Side:          input.Side,
			AccountItemID: input.AccountItemID,
			Amount:        input.Amount,
			PartnerID:     input.PartnerID,
			CreatedAt:     now,
		}
		next.Lines = append(next.Lines, line)
	}

	if debitTotal != creditTotal {
		return reversal, nil, fmt.Errorf("debit and credit totals do not match: %f != %f", debitTotal, creditTotal)
	}

	if err := u.journalRepo.Save(next); err != nil {
		return reversal, nil, fmt.Errorf("failed to save new entry: %w", err)
	}

	return reversal, next, nil
}
