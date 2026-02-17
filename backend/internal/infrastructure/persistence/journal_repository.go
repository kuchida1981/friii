package persistence

import (
	"database/sql"
	"time"

	"github.com/kuchida1981/friii/internal/domain/entity"
	"github.com/kuchida1981/friii/internal/domain/repository"

	"github.com/google/uuid"
)

type journalRepository struct {
	db *sql.DB
}

func NewJournalRepository(db *sql.DB) repository.JournalRepository {
	return &journalRepository{db: db}
}

func (r *journalRepository) Save(entry *entity.JournalEntry) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// 1. ヘッダーの保存
	headerQuery := `
		INSERT INTO journal_entries (id, transaction_date, description, original_entry_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.Exec(headerQuery, entry.ID, entry.TransactionDate, entry.Description, entry.OriginalEntryID, entry.CreatedAt)
	if err != nil {
		return err
	}

	// 2. 明細の保存
	lineQuery := `
		INSERT INTO journal_lines (id, entry_id, side, account_item_id, amount, partner_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	for _, line := range entry.Lines {
		_, err = tx.Exec(lineQuery, line.ID, entry.ID, line.Side, line.AccountItemID, line.Amount, line.PartnerID, line.CreatedAt)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *journalRepository) FindByID(id uuid.UUID) (*entity.JournalEntry, error) {
	entryQuery := `SELECT id, transaction_date, description, original_entry_id, created_at FROM journal_entries WHERE id = $1`
	entry := &entity.JournalEntry{}
	err := r.db.QueryRow(entryQuery, id).Scan(
		&entry.ID, &entry.TransactionDate, &entry.Description, &entry.OriginalEntryID, &entry.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if err := r.loadLines(entry); err != nil {
		return nil, err
	}

	return entry, nil
}

func (r *journalRepository) List(from, to time.Time) ([]*entity.JournalEntry, error) {
	query := `
		SELECT id, transaction_date, description, original_entry_id, created_at 
		FROM journal_entries 
		WHERE transaction_date BETWEEN $1 AND $2
		ORDER BY transaction_date ASC, created_at ASC
	`
	rows, err := r.db.Query(query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*entity.JournalEntry
	for rows.Next() {
		entry := &entity.JournalEntry{}
		if err := rows.Scan(&entry.ID, &entry.TransactionDate, &entry.Description, &entry.OriginalEntryID, &entry.CreatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// 各エントリーの明細を取得
	for _, entry := range entries {
		if err := r.loadLines(entry); err != nil {
			return nil, err
		}
	}

	return entries, nil
}

func (r *journalRepository) loadLines(entry *entity.JournalEntry) error {
	lineQuery := `SELECT id, entry_id, side, account_item_id, amount, partner_id, created_at FROM journal_lines WHERE entry_id = $1`
	rows, err := r.db.Query(lineQuery, entry.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		line := &entity.JournalLine{}
		if err := rows.Scan(&line.ID, &line.EntryID, &line.Side, &line.AccountItemID, &line.Amount, &line.PartnerID, &line.CreatedAt); err != nil {
			return err
		}
		entry.Lines = append(entry.Lines, line)
	}
	return rows.Err()
}
