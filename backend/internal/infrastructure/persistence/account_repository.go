package persistence

import (
	"database/sql"
	"time"

	"github.com/kuchida1981/friii/internal/domain/entity"
	"github.com/kuchida1981/friii/internal/domain/repository"

	"github.com/google/uuid"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) repository.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) SaveItem(item *entity.AccountItem) error {
	query := `
		INSERT INTO account_items (id, category_id, code, name, valid_from, valid_to, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, item.ID, item.CategoryID, item.Code, item.Name, item.ValidFrom, item.ValidTo, item.CreatedAt)
	return err
}

func (r *accountRepository) FindItemByID(id uuid.UUID, at time.Time) (*entity.AccountItem, error) {
	query := `
		SELECT a.id, a.category_id, a.code, a.name, a.valid_from, a.valid_to, a.created_at,
		       c.id, c.name, c.normal_side, c.report_type
		FROM account_items a
		JOIN account_categories c ON a.category_id = c.id
		WHERE a.id = $1 AND a.valid_from <= $2 AND (a.valid_to IS NULL OR a.valid_to > $2)
	`
	item := &entity.AccountItem{Category: &entity.AccountCategory{}}
	err := r.db.QueryRow(query, id, at).Scan(
		&item.ID, &item.CategoryID, &item.Code, &item.Name, &item.ValidFrom, &item.ValidTo, &item.CreatedAt,
		&item.Category.ID, &item.Category.Name, &item.Category.NormalSide, &item.Category.ReportType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r *accountRepository) ListItem(at time.Time) ([]*entity.AccountItem, error) {
	query := `
		SELECT a.id, a.category_id, a.code, a.name, a.valid_from, a.valid_to, a.created_at,
		       c.id, c.name, c.normal_side, c.report_type
		FROM account_items a
		JOIN account_categories c ON a.category_id = c.id
		WHERE a.valid_from <= $1 AND (a.valid_to IS NULL OR a.valid_to > $1)
		ORDER BY a.code ASC
	`
	rows, err := r.db.Query(query, at)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*entity.AccountItem
	for rows.Next() {
		item := &entity.AccountItem{Category: &entity.AccountCategory{}}
		err := rows.Scan(
			&item.ID, &item.CategoryID, &item.Code, &item.Name, &item.ValidFrom, &item.ValidTo, &item.CreatedAt,
			&item.Category.ID, &item.Category.Name, &item.Category.NormalSide, &item.Category.ReportType,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *accountRepository) FindCategoryByID(id uuid.UUID) (*entity.AccountCategory, error) {
	query := `SELECT id, name, normal_side, report_type FROM account_categories WHERE id = $1`
	c := &entity.AccountCategory{}
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.NormalSide, &c.ReportType)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *accountRepository) ListCategories() ([]*entity.AccountCategory, error) {
	query := `SELECT id, name, normal_side, report_type FROM account_categories ORDER BY id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.AccountCategory
	for rows.Next() {
		c := &entity.AccountCategory{}
		if err := rows.Scan(&c.ID, &c.Name, &c.NormalSide, &c.ReportType); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
