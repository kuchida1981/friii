package persistence

import (
	"database/sql"
	"time"

	"github.com/kuchida1981/friii/internal/domain/entity"
	"github.com/kuchida1981/friii/internal/domain/repository"

	"github.com/google/uuid"
)

type partnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) repository.PartnerRepository {
	return &partnerRepository{db: db}
}

func (r *partnerRepository) Save(p *entity.Partner) error {
	query := `
		INSERT INTO partners (id, name, registration_number, valid_from, valid_to, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, p.ID, p.Name, p.RegistrationNumber, p.ValidFrom, p.ValidTo, p.CreatedAt)
	return err
}

func (r *partnerRepository) FindByID(id uuid.UUID, at time.Time) (*entity.Partner, error) {
	query := `
		SELECT id, name, registration_number, valid_from, valid_to, created_at
		FROM partners
		WHERE id = $1 AND valid_from <= $2 AND (valid_to IS NULL OR valid_to > $2)
	`
	p := &entity.Partner{}
	err := r.db.QueryRow(query, id, at).Scan(
		&p.ID, &p.Name, &p.RegistrationNumber, &p.ValidFrom, &p.ValidTo, &p.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *partnerRepository) List(at time.Time) ([]*entity.Partner, error) {
	query := `
		SELECT id, name, registration_number, valid_from, valid_to, created_at
		FROM partners
		WHERE valid_from <= $1 AND (valid_to IS NULL OR valid_to > $1)
		ORDER BY name ASC
	`
	rows, err := r.db.Query(query, at)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []*entity.Partner
	for rows.Next() {
		p := &entity.Partner{}
		if err := rows.Scan(&p.ID, &p.Name, &p.RegistrationNumber, &p.ValidFrom, &p.ValidTo, &p.CreatedAt); err != nil {
			return nil, err
		}
		partners = append(partners, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return partners, nil
}
